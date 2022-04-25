// パッケージ宣言は一つだけでないとエラーが出る
package main

// パッケージのインポート
import (
	"fmt"
	"time"
)

func main() {
	// Hello world出力
	fmt.Println("Hello World")

	fmt.Println(time.Now())

	// ターミナルで[go run main.go]
}