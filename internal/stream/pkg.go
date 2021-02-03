package stream

import (
	"io"
	"io/ioutil"

	"go.octolab.org/async"
	"go.octolab.org/safe"
	"go.octolab.org/unsafe"
)

type Operator interface {
	Operate() error
}

type Pipeline interface {
	Pipe(...func(Reader, Writer) Operator) Operator
}

type Valve = func(Reader, Writer) Operator

type OperatorFunc func() error

func (fn OperatorFunc) Operate() error {
	return fn()
}

type PipelineFunc func(...Valve) Operator

func (fn PipelineFunc) Pipe(valves ...Valve) Operator {
	return fn(valves...)
}

func Connect(input Reader, output Writer) Pipeline {
	return PipelineFunc(func(valves ...Valve) Operator {
		return OperatorFunc(func() error {
			job := new(async.Job)
			defer job.Wait()

			for _, build := range valves {
				in, out := io.Pipe()
				operator := build(input, out)
				job.Do(func() error {
					defer safe.Close(out, unsafe.Ignore)
					return operator.Operate()
				}, Discard(in))
				input = in
			}

			job.Do(Copy(input, output).Operate, Discard(input))

			return nil
		})
	})
}

func Copy(input Reader, output Writer) Operator {
	return OperatorFunc(func() error {
		_, err := io.Copy(output, input)
		return err
	})
}

func Discard(input Reader) func(error) {
	return func(error) {
		unsafe.DoSilent(io.Copy(ioutil.Discard, input))
	}
}
