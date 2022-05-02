package main

import (
	"fmt"
	// "golang_practice/section13/foo"
)

// スコープ

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s
	b = s
	return s
}

func main() {

	// fmt.Println(foo.Max)

	fmt.Println(appName())

	fmt.Println(Do("AAA"))

}