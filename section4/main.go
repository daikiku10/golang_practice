package main

import "fmt"

// 型

func main() {
	
	/** 整数型 */
	var i int = 100

	fmt.Println(i + 50)

	// 型を調べる方法
	fmt.Printf("%T\n", i)

	// 型変換
	fmt.Printf("%T\n", int32(i))

	/** 不動小数点型　*/
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	/** 論理値型 */
	var t, f bool = true, false
	fmt.Println(t, f)

	/** 文字列型 */
	var s string = "test"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	fmt.Println(`aaa
	aaa
	     aa`)

	/** byte(unit8)型 */
	byteA := []byte{72,73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	/** 配列型 */
	var arr1 [3]int
	fmt.Println(arr1)

	var arr2 [3]string = [3]string{"a", "b", "c"}
	fmt.Println(arr2)

	arr3 := [3]int{1,2,3}
	fmt.Println(arr3)

	arr4 := [...]string{"a", "b"}
	fmt.Println(arr4)

	/** interface nil */
	var x interface{}
	fmt.Println(x)
}