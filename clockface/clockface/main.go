package main

import (
	"os"
	"time"

	"github.com/tsugoshi/learn-go-with-tests/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
