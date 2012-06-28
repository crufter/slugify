package main

import(
	"github.com/opesun/slugify"
	"fmt"
)

func main() {
	str := `This is some hungarian text with accents: "Helló, belló."`
	fmt.Println(slugify.S(str))
}