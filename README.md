# Minions

A golang library to help you manage worker pools


# Usage

```go
package main

import (
	"log"
	"time"

	"github.com/eteissonniere/minions"
)

func main() {
	log.Println("starting")

	nbMinions := 10
	pool := minions.Hire(nbMinions)

	for i := 0; i<nbMinions*2; i++ {
		// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		i := i
		pool.Execute(minions.Command{
			Instructions: func() {
				time.Sleep(time.Duration(nbMinions - i) * time.Millisecond)
				log.Println(i)
			},
		})
	}

	pool.Terminate()
}
```