package main

import "fmt"

/** 関数 */

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func NoReturn() {
	fmt.Println("No Return")
}

func Double(price int) (result int) { // 返り値の変数を指定している。そのため、returnだけで良い。
	result = price * 2
	return
}

/** 関数を返す関数 */
func ReturnFunc() func() {
	return func() {
		fmt.Println("私は関数である。")
	}
}

/** 関数を引数に取る関数 */
func CallFunction(f func()) {
	f()
}

/** クロージャーの実装 */
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

/** ジェネレーターの実装 */
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3)

	i4 := Double(1000)
	fmt.Println(i4)

	NoReturn()

	/** 無名関数 */
	f := func(x, y int) int {
		return x + y
	}

	r := f(1, 2)

	fmt.Println(r)

	d := ReturnFunc()
	d()

	/** 関数を引数に取る関数 */
	CallFunction(func() {
		fmt.Println("私は関数である。")
	})

	/** クロージャーの実装 */
	fa := Later()
	fmt.Println(fa("Hello"))

	/** ジェネレーターの実装 */
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

}