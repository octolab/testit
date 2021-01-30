module testdata

go 1.15

// go test -count=1 -run ^Fake$$ ./... | pbcopy
// go test -run ^Fake$$ ./... | pbcopy
// go test -run ^Fake$$ ./... |& pbcopy
//
// go test -count=1 ./... | pbcopy
// go test ./... | pbcopy
// go test ./... |& pbcopy
