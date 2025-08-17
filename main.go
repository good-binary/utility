package main

import (
	"fmt"
	"time"

	"github.com/good-binary/utility/random"
)

func main() {
	fmt.Println(random.RandomFullName())
	t := time.Now()
	fmt.Println(t.Format("20060102150405"))
}
