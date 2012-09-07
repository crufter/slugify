package main

import (
	"fmt"
	"github.com/opesun/slugify"
)

func main() {
	str := `This is some hungarian text with accents: "Helló, belló."`
	fmt.Println(slugify.S(str))
}
