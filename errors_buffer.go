package liberrors

import (
	"fmt"
	"io"
	"strings"

	libescapes "github.com/bbfh-dev/lib-ansi-escapes"
)

type Buffer struct {
	FirstLine   uint
	Buffer      string
	Highlighted string
}

func (buffer Buffer) IsEmpty() bool {
	return len(buffer.Buffer) == 0 && len(buffer.Highlighted) == 0
}

func (buffer Buffer) Print(writer io.Writer) {
	writer.Write([]byte(libescapes.TextColorWhite + "\n"))
	var line_index uint

	for line := range strings.SplitSeq(buffer.Buffer, "\n") {
		if line_index != 0 {
			writer.Write([]byte{'\n'})
		}
		fmt.Fprintf(
			writer,
			"%5d |  %s",
			buffer.FirstLine+line_index,
			line,
		)
		line_index++
	}

	writer.Write([]byte(libescapes.TextColorBrightRed))

	if len(buffer.Highlighted) == 0 {
		writer.Write([]byte("←—"))
	} else {
		i := uint(0)
		contents := strings.TrimSuffix(buffer.Highlighted, "\n")
		for line := range strings.SplitSeq(contents, "\n") {
			if i == 0 {
				writer.Write([]byte(line))
			} else {
				fmt.Fprintf(
					writer,
					"\n%5d |  %s",
					buffer.FirstLine+line_index+i-1,
					line,
				)
			}
			i++
		}
	}

	writer.Write([]byte(libescapes.ColorReset + "\n"))
}
