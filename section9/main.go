package main

import "fmt"


func main() {
	// スライス (配列と違って可変長であるため、要素追加などのサイズ変更ができる。)
	// 宣言、操作
	var slice []int = []int{100, 200}
	fmt.Println(slice)

	// append スライスに要素を追加する。
	slice = append(slice, 300)
	fmt.Println(slice)

	// slice[0] = 1000
	// fmt.Println(slice)

	// 暗黙的なスライス
	slice2 := []string{"a", "b"}
	fmt.Println(slice2)

	// make スライスを生成する。第二引数は要素数、第三引数は容量を指定できる。
	slice3 := make([]int, 5, 10)
	fmt.Println(slice3)

	// len関数 スライスの数を調べる
	fmt.Println(len(slice3))

	// cpa関数 スライスの容量を調べる
  fmt.Println(cap(slice3))

	// copy
	copySlice := make([]int, 2)
	// nは成功した時の要素数が入る。
	n := copy(copySlice, slice)
	fmt.Println(n, copySlice)

	// マップ
	var mapData = map[string]int{"A": 100, "B": 200}
	fmt.Println(mapData)

	// channel
	// 複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造。
	var channel chan int

	channel = make(chan int)
	fmt.Println(channel)

	ch3 := make(chan int, 5)
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	fmt.Println(len(ch3))

	i := <- ch3
	fmt.Println(i)
	
	i2 := <- ch3
	fmt.Println(i2)

	// 受信専用
	// var channel2 <-chan int

	// 送信専用
	// var channel3 chan<- int

	// チャネルとゴルーチン
	ch1 := make(chan int)

	fmt.Println(<- ch1)

}