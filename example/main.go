package main

import (
	"github.com/go-steven/cerr"
	"fmt"
)

func main() {
	err := cerr.New("test error")
	fmt.Printf("err : %v\n", err.Error())

	err = cerr.NewCode(100, "test error")
	fmt.Printf("err : %v\n", err.Error())

	err = cerr.Errorf("found error: %v", "test error")
	fmt.Printf("err : %v\n", err.Error())

	err = cerr.ErrorfCode(100, "found error: %v", "test error")
	fmt.Printf("err : %v\n", err.Error())
}

