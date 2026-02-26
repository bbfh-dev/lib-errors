package liberrors_test

import (
	"testing"

	liberrors "github.com/bbfh-dev/lib-errors"
)

func TestFormatDir(t *testing.T) {
	derr := &liberrors.DetailedError{
		Label: liberrors.ERR_INTERNAL,
		Context: liberrors.DirContext{
			Path: "/tmp/some/path/to/a/dir",
		},
		Details: "This is an example error body",
	}
	derr.Print(t.Output())
}

func TestFormatNoContext(t *testing.T) {
	derr := &liberrors.DetailedError{
		Label:   liberrors.ERR_INTERNAL,
		Context: nil,
		Details: "This is an error without any context",
	}
	derr.Print(t.Output())
}

func TestFormatFileNoBuffer(t *testing.T) {
	derr := &liberrors.DetailedError{
		Label: liberrors.ERR_INTERNAL,
		Context: liberrors.FileContext{
			Trace: []liberrors.TraceItem{
				{
					Name: "/tmp/some/path/to/a/file.txt",
					Row:  6,
					Col:  9,
				},
			},
			Buffer: liberrors.Buffer{},
		},
		Details: "This is an example error body",
	}
	derr.Print(t.Output())
}

func TestFormatFile(t *testing.T) {
	derr := &liberrors.DetailedError{
		Label: liberrors.ERR_INTERNAL,
		Context: liberrors.FileContext{
			Trace: []liberrors.TraceItem{
				{
					Name: "nested_file.txt",
					Row:  0,
					Col:  2,
				},
				{
					Name: "/tmp/some/path/to/a/file.txt",
					Row:  6,
					Col:  9,
				},
			},
			Buffer: liberrors.Buffer{
				FirstLine:   5,
				Buffer:      "Hello World\ntesting ",
				Highlighted: "123",
			},
		},
		Details: "This is an example error body",
	}
	derr.Print(t.Output())
}

func TestFormatFileNoHighlighted(t *testing.T) {
	derr := &liberrors.DetailedError{
		Label: liberrors.ERR_INTERNAL,
		Context: liberrors.FileContext{
			Trace: []liberrors.TraceItem{
				{
					Name: "/tmp/some/path/to/a/file.txt",
					Row:  6,
					Col:  9,
				},
			},
			Buffer: liberrors.Buffer{
				FirstLine:   5,
				Buffer:      "Hello World\ntesting ",
				Highlighted: "",
			},
		},
		Details: "This is an example error body",
	}
	derr.Print(t.Output())
}

func TestFormatProgram(t *testing.T) {
	derr := &liberrors.DetailedError{
		Label: liberrors.ERR_INTERNAL,
		Context: liberrors.ProgramContext{
			Binary: "/bin/ls",
			Args:   []string{"-l", "/tmp"},
			Stderr: "ls: /tmp: no such file or directory",
		},
		Details: "This is an example error body",
	}
	derr.Print(t.Output())
}
