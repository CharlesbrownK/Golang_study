# Print Function


## What is "Print"

 The `Print()` function prints the specified message to the screen, or other standard output device.

The message can be a string, or any other object, the object will be converted into a string before written to the screen.


## Difference between "Printf" and "Println"


### Printf

 `Printf` function is short for print formatter, it gives you the ability to mark where in the String variables will go and pass in those variables with it.

Try this code in your editor!

```go
package main

import (
	"fmt"
)

func main() {
	var j float32 = 27
	k := 99

	fmt.Printf("%v, %T", k, k)
	fmt.Printf("%f, %T", j, j)
}
```

The printed result will be this:

```
99, int27.000000, float32
```


### Println

 `Println` is short for "print line", meaning after the argument is printed then goes to the next line.

Try this code in your editor!

```go
package main

import (
	"fmt"
)

var (
	AuthorName          string = "Charlesbrown K"
	AuthorGithubAddress string = "https://github.com/CharlesbrownK"
	AuthorPageAddress   string = "https://charlesbrownk.github.io/"
)

func main() {
	fmt.Println(AuthorName)
	fmt.Println(AuthorGithubAddress, AuthorPageAddress)
}
```

The printed result will be this:

```
Charlesbrown K
https://github.com/CharlesbrownK https://charlesbrownk.github.io/
```