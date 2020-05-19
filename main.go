package main

import (
	"fmt"

	"github.com/nooncall/ci-base/pkg/utils"
)

func main() {
	fmt.Println(utils.FooWord("hello world!"))
	fmt.Println("do something bad!")
}
