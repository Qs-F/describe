package main

import (
	"fmt"
	"github.com/Qs-F/coloring"
)

func main() {
	fmt.Printf(`
./
%s

./Hatena
%s

./Hatena/bookmark
%s

./Google+/Described/Fumi
%s

./Tw
%s
`, coloring.Yellow("Test Directory"), coloring.Yellow("Wow, this is proj.\nHellooo."), coloring.Yellow("Wow wow"), coloring.Yellow("Long\nLong\nLong\nago, in the galaxy far away..."), coloring.Yellow("Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem Lorem "))
}
