package main

import (
	"fmt"
	"github.com/truxcoder/trux/cmd/trux"
)

func main() {
	err := trux.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
