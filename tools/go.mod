module go.octolab.org/toolset/testit/tools

go 1.15

require (
	github.com/golang/mock v1.4.4
	github.com/golangci/golangci-lint v1.36.0
	github.com/kyoh86/looppointer v0.1.7
	github.com/maruel/panicparse/v2 v2.1.1
	github.com/rakyll/gotest v0.0.5
	golang.org/x/exp v0.0.0-20210201131500-d352d2db2ceb
	golang.org/x/tools v0.1.0
)

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.1.0
