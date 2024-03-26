package main

import "fmt"

// 以下の要件を満たす定数を定義してください
// 1. 色を表す定数を定義する
//    - RED, GREEN, BLUE, YELLOW, ORANGE, PURPLE の6つの色を定義する
//    - 値は iota を使って割り当てる
// 2. サイズを表す定数を定義する
//    - SMALL, MEDIUM, LARGE の3つのサイズを定義する
//    - 値は iota を使って割り当てる
//    - SMALL の値は 1 から始まるようにする

// ここで定数を定義する
const (
	RED = iota
	GREEN
	BLUE
	YELLOW
	ORANGE
	PURPLE
)

const (
	_ = iota
	SMALL
	MEDIUM
	LARGE
)

func main() {
	// 以下の出力になるように fmt.Println を使って出力する
	// RED: 0
	// GREEN: 1
	// BLUE: 2
	// YELLOW: 3
	// ORANGE: 4
	// PURPLE: 5
	// 色の定数の値を出力する
	fmt.Println("RED:", RED)
	fmt.Println("GREEN:", GREEN)
	fmt.Println("BLUE:", BLUE)
	fmt.Println("YELLOW:", YELLOW)
	fmt.Println("ORANGE:", ORANGE)
	fmt.Println("PURPLE:", PURPLE)
	// SMALL: 1
	// MEDIUM: 2
	// LARGE: 3
	fmt.Println("SMALL:", SMALL)
	fmt.Println("MEDIUM:", MEDIUM)
	fmt.Println("LARGE:", LARGE)

	// ここで出力処理を書く
}
