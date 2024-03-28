package main

import "fmt"

func main() {
	// 1. string型のキーとint型の値を持つmapを作成し、適当なキーと値のペアを3つ追加する
	m := map[string]int{
		"a":   1,
		"ab":  2,
		"abc": 3,
	}

	// 2. mapの各キーと値のペアを出力する
	for key, value := range m {
		fmt.Println(key, value)
	}

	// 3. int型の要素を持つ長さ5の配列を作成し、適当な値で初期化する
	arr := [5]int{1, 2, 3, 4, 5}

	// 4. 配列の各要素を2倍した値に更新する
	for i, v := range arr {
		arr[i] = v * 2
	}

	// 5. 配列の要素を出力する

	fmt.Println(arr)

	// 6. 手順3で作成した配列をもとにスライスを作成する

	// ここでスライスを作成する

	// 7. スライスの要素を出力する

	// ここでスライスの要素を出力する

	// 8. スライスの容量を出力する

	// ここでスライスの容量を出力する

	// 9. スライスに新しい要素を追加する

	// ここでスライスに新しい要素を追加する

	// 10. スライスの要素を出力する

	// ここでスライスの要素を出力する

	// 11. スライスの容量を出力する

	// ここでスライスの容量を出力する
}