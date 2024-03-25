package main

import "fmt"

func main() {
	// 1. 整数のスライスnumbersを定義し、適当な値で初期化する
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}

	// 2. findMostFrequent関数を呼び出し、最頻値を求める
	mostFrequent := findMostFrequent(numbers)

	// 3. 最頻値を出力する
	fmt.Println("最頻値:", mostFrequent)
}

// findMostFrequent関数を定義する
// 引数:
//   - numbers: 整数のスライス
//
// 戻り値:
//   - 最頻値の整数
func findMostFrequent(numbers []int) int {
	// ここに処理を書く
}

//期待する出力:
//最頻値: 3
