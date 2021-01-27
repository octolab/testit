module tool/tools

go 1.15

require (
	github.com/golang/mock v1.4.4
	github.com/golangci/golangci-lint v1.36.0
	github.com/kyoh86/looppointer v0.1.7
	golang.org/x/exp v0.0.0-20210126160131-4ebcfff34a87
	golang.org/x/tools v0.1.0
)

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.1.0
