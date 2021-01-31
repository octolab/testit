package stream

import (
	"bufio"
	"fmt"
	"strings"
)

type Processor interface {
	Process() error
}

type Process func() error

func (fn Process) Process() error {
	return fn()
}

func GoTestCompile(input Reader, output Writer) Processor {
	return Process(func() error {
		scanner := bufio.NewScanner(input)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()

			if strings.Contains(line, "no test files") {
				continue
			}
			if strings.Contains(line, "no tests to run") {
				continue
			}

			if _, err := fmt.Fprintln(output, line); err != nil {
				return err
			}
		}
		return nil
	})
}
