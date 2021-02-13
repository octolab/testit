// Package failed has table tests with failed assertion.
package failed

func Concat(strings ...string) string {
	var result string
	for _, str := range strings {
		result += str
	}
	return result
}
