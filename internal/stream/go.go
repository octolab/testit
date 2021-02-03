package stream

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func GoTest(input Reader, output Writer) Operator {
	const (
		fail = color.FgHiRed
		pass = color.FgGreen
		skip = color.FgYellow
	)

	return OperatorFunc(func() error {
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

func GoTestCompile(input Reader, output Writer) Operator {
	return OperatorFunc(func() error {
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
