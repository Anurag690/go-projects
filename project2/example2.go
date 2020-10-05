package main

import (
	"fmt"

	"github.com/Anurag690/go-projects/project1"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello example2!")
	project1.DoSomething()
	cmp.Diff("f", "h")
}
