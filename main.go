package main

import (
	"fmt"

	"github.com/kloud-warrior/go-beginner-01/queue"
)

func main() {
	q := queue.New()
	fmt.Println(q.IsEmpty())
}
