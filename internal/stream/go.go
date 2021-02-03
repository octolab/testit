package stream

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
	"go.octolab.org/async"
	"go.octolab.org/safe"
	"go.octolab.org/unsafe"
)

type Processor interface {
	Process() error
}

type Process func() error

func (fn Process) Process() error {
	return fn()
}

func Copy(input Reader, output Writer) Processor {
	return Process(func() error {
		_, err := io.Copy(output, input)
		return err
	})
}

func Discard(input Reader) func(error) {
	return func(err error) {
		unsafe.DoSilent(io.Copy(ioutil.Discard, input))
	}
}

func Pipe(input Reader, output Writer, processors ...func(Reader, Writer) Processor) Processor {
	return Process(func() error {
		job := new(async.Job)
		defer job.Wait()

		for _, build := range processors {
			in, out := io.Pipe()
			processor := build(input, out)
			job.Do(func() error {
				defer safe.Close(out, unsafe.Ignore)
				return processor.Process()
			}, Discard(in))
			input = in
		}

		job.Do(Copy(input, output).Process, Discard(input))

		return nil
	})
}

func GoTest(input Reader, output Writer) Processor {
	const (
		fail = color.FgHiRed
		pass = color.FgGreen
		skip = color.FgYellow
	)

	return Process(func() error {
		scanner := bufio.NewScanner(input)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()

			var class color.Attribute
			trimmed := strings.TrimSpace(line)
			switch {
			case strings.HasPrefix(trimmed, "--- FAIL"):
				fallthrough
			case strings.HasPrefix(trimmed, "FAIL"):
				class = fail

			case strings.HasPrefix(trimmed, "--- PASS"):
				fallthrough
			case strings.HasPrefix(trimmed, "PASS"):
				fallthrough
			case strings.HasPrefix(trimmed, "ok"):
				class = pass

			case strings.Contains(trimmed, noTestFiles):
				fallthrough
			case strings.Contains(trimmed, noTestsToRun):
				fallthrough
			case strings.HasPrefix(trimmed, "--- SKIP"):
				class = skip
			}

			var printer = fmt.Fprintln
			if class > 0 {
				printer = color.New(class).Fprintln
			}
			if _, err := printer(output, line); err != nil {
				return err
			}
		}
		return nil
	})
}

func GoTestCompile(input Reader, output Writer) Processor {
	return Process(func() error {
		scanner := bufio.NewScanner(input)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()

			if strings.Contains(line, noTestFiles) {
				continue
			}
			if strings.Contains(line, noTestsToRun) {
				continue
			}

			if _, err := fmt.Fprintln(output, line); err != nil {
				return err
			}
		}
		return nil
	})
}

const (
	noTestFiles  = "[no test files]"
	noTestsToRun = "[no tests to run]"
)
