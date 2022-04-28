package main

import "fmt"


func main() {
	/** 算術演算子 */
	fmt.Println(1 + 2)

	fmt.Println("テスト" + "太郎")

	fmt.Println(5 - 2)

	fmt.Println(2 * 2)

	fmt.Println(6 / 2)

	fmt.Println(6 % 5)

	/** 比較演算子 */
	fmt.Println(1 == 1)

	fmt.Println(1 == 2)

	fmt.Println(4 <= 8)

	fmt.Println(4 <= 3)

	fmt.Println(4 < 4)

	/** 論理演算子 */

	fmt.Println(true && false == true)

	fmt.Println(true && 1 == 1)

	fmt.Println(true || 2 == 1)

}