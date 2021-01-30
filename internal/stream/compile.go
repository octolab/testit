package stream

import (
	"bufio"
	"fmt"
	"strings"
)

func GoCompileProcess(input Reader, output Writer) error {
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
}
