package main

import "fmt"

// defer
func TestDefer() {
	// 関数が終了するときにdeferで登録したプログラムを実行する。
	defer fmt.Println("END")
	fmt.Println("START")
}

func main() {
	// if
	// 条件分岐
	a := 2
	if a == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("Not")
	}

	// 簡易文付きif文
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	// for文
	// 繰り返し処理
	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	var x interface{} = 100

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is string")
	} else {
		fmt.Println("i don't know")
	}

	// switch 型スイッチ
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	// defer
	TestDefer()

	defer func() {
		defer fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
}