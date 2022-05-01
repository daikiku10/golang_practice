package main

import "fmt"

func main() {
	// ポインタ
	var n int = 100
	fmt.Println(n)
	// アドレス表示
	fmt.Println(&n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	*p = 300
	fmt.Println(n)

}