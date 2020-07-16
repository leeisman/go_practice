package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// unbuffered
	fmt.Fprintf(os.Stdout, "HTTP/1.1 200 Connection established")
	// buffered: os.Stdout implements io.Writer
	buf := bufio.NewWriter(os.Stdout)
	// and now so does buf.
	fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
	buf.Flush()
}
