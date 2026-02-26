package liberrors

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	libescapes "github.com/bbfh-dev/lib-ansi-escapes"
)

type Context interface {
	Print(io.Writer)
}

// ————————————————————————

type TraceItem struct {
	Name string
	// Use negative values to hide it from the output
	Col, Row int
}

func (item TraceItem) PrintRoot(writer io.Writer) {
	fmt.Fprintf(writer, "    in %s", item.Name)
	item.printLocation(writer)
}

func (item TraceItem) PrintNested(writer io.Writer) {
	fmt.Fprintf(writer, "    └─ from %s", item.Name)
	item.printLocation(writer)
}

func (item TraceItem) printLocation(writer io.Writer) {
	if item.Row >= 0 {
		fmt.Fprintf(writer, ":%d", item.Row)
	}
	if item.Col >= 0 {
		fmt.Fprintf(writer, ":%d", item.Col)
	}
	writer.Write([]byte{'\n'})
}

// ————————————————————————

type DirContext struct {
	Path string
}

func (context DirContext) Print(writer io.Writer) {
	item := TraceItem{Name: context.Path, Row: -1, Col: -1}
	item.PrintRoot(writer)
}

func NewDirContext(path string) DirContext {
	return DirContext{Path: path}
}

type FileContext struct {
	Trace  []TraceItem
	Buffer Buffer
}

func (context FileContext) Print(writer io.Writer) {
	if len(context.Trace) > 0 {
		context.Trace[0].PrintRoot(writer)
		for _, item := range context.Trace[1:] {
			item.PrintNested(writer)
		}
	}

	if !context.Buffer.IsEmpty() {
		context.Buffer.Print(writer)
	}
}

type ProgramContext struct {
	Binary string
	Args   []string
	Stderr string
}

func (context ProgramContext) Print(writer io.Writer) {
	writer.Write([]byte(libescapes.TextColorBrightGreen))
	fmt.Fprintf(writer, "    $ %s %s", context.Binary, strings.Join(context.Args, " "))
	writer.Write([]byte(libescapes.ColorReset + "\n"))

	writer.Write([]byte(libescapes.TextColorWhite))
	fmt.Fprint(writer, "╾─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─\n")

	writer.Write([]byte(libescapes.TextColorBrightRed))
	fmt.Fprintf(writer, "%s\n", context.Stderr)

	writer.Write([]byte(libescapes.TextColorWhite))
	fmt.Fprint(writer, "╾─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─")
	writer.Write([]byte(libescapes.ColorReset))
}

func NewProgramContext(cmd exec.Cmd, stderr string) ProgramContext {
	return ProgramContext{
		Binary: cmd.Path,
		Args:   cmd.Args[1:],
		Stderr: stderr,
	}
}
