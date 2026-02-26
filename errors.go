// Detailed, pretty-printed colorful errors for Go
package liberrors

import (
	"fmt"
	"io"

	libescapes "github.com/bbfh-dev/lib-ansi-escapes"
)

// Suggested values for DetailedError.Label
const (
	ERR_INTERNAL = "Internal Error"
	ERR_ASSERT   = "Assertion Error"
	ERR_IO       = "I/O Error"
	ERR_READ     = "Reading Error"
	ERR_WRITE    = "Writing Error"
	ERR_SYNTAX   = "Syntax Error"
	ERR_FORMAT   = "Formatting Error"
	ERR_VALIDATE = "Validation Error"
	ERR_EXECUTE  = "Execution Error"
	ERR_CONVERT  = "Conversion Error"
)

// An improved pretty-printed detailed error with context
type DetailedError struct {
	Label   string
	Context Context
	Details string
}

func (err *DetailedError) Print(writer io.Writer) {
	writer.Write([]byte(libescapes.TextColorWhite + "────────────────────────────────\n"))
	writer.Write([]byte(" [!] " + libescapes.TextColorBrightRed + err.Label +
		"\n" + libescapes.TextColorWhite))

	if err.Context != nil {
		err.Context.Print(writer)
	}

	writer.Write([]byte(libescapes.TextColorBrightYellow +
		"\n >>> " + libescapes.ColorReset + err.Details + "\n"))
}

func (err *DetailedError) Error() string {
	return fmt.Sprintf("(%s) %s", err.Label, err.Details)
}

// Constructor for the most common error
func NewIO(err error, path string) error {
	if err == nil {
		return nil
	}
	return &DetailedError{
		Label:   ERR_IO,
		Context: NewDirContext(path),
		Details: err.Error(),
	}
}

// Prints DetailedError{} or generic error to writer
//
// Example usage:
//
//	libparsex.Print(err, os.Stderr)
func Print(err error, writer io.Writer) {
	switch err := err.(type) {
	case *DetailedError:
		err.Print(writer)
	default:
		io.WriteString(writer, err.Error())
	}

	io.WriteString(writer, "\n")
}
