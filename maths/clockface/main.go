package main

import (
	. "go-book/maths/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()
	Write(os.Stdout, t)
}
