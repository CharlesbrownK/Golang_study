package main

import (
	"fmt"
)

var (
	AuthorName        string = "Charlesbrown K"
	GithubAddress     string = "https://github.com/CharlesbrownK"
	AuthorPageAddress string = "https://charlesbrownk.github.io/"
	AuthorAge         int    = 18
)

var (
	counter float32 = 0
)

func main() {
	fmt.Printf(AuthorName)
	fmt.Printf("\n%v, %T", counter, counter)
}
