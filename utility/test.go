package main

import (
	"fmt"
	"io/ioutil"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func main() {
	fmt.Println("vim-go")
	bytes, _ := ioutil.ReadFile("test.md")
	rendered := (blackfriday.Run(bytes))
	fmt.Println(string(rendered))

}
