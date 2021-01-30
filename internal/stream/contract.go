package stream

import "io"

//go:generate mockgen -source $GOFILE -destination mocks_test.go -package ${GOPACKAGE}_test

type Reader interface {
	io.Reader
}

type Writer interface {
	io.Writer
}
