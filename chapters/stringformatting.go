package chapters

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var f = fmt.Printf

func StringFormatting() {
	p := point{1, 3}

	f("struct: %vn", p)

	f("struct2: %+v\n", p)

	f("struct3: %#v\n", p)

	f("type: %T\n", p)

	f("bool: %t\n", true)

	f("bin: %b\n", 123)

	f("char: %c\n", 33)

	f("hex: %x\n", 546)

	f("float: %f\n", 78.9)

	f("float2: %e\n", 123400000.0)

	f("float3: %E\n", 1234000000.0)

	f("width1: |%6d|%6d|\n", 12, 123)

	f("width2: |%6.2f|%6.3f|\n", 1.2, 3.45)

	f("width2: |%-6.2f|%-6.3f|\n", 1.2, 3.45)

	f("width2: |%6s|%6s|\n", "foo", "bar")

	f("width2: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an%s\n", "error")
}
