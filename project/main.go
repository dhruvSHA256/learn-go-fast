package main

import (
	// "errors"
	"fmt"
	// "os"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println(errors.New("missing arsg").Error())
	// 	os.Exit(1)
	// }
	// name := os.Args[1]
	// fmt.Printf("hello %v\n", name)
	// var colors map[string]string
	colors := make(map[string]string)
	// colors := map[string]string{
	// 	"a": "b",
	// }
	colors["a"] = "b"
	delete(colors, "a")
	fmt.Printf("%+v", colors)

}
