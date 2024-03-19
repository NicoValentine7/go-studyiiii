package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// 1. 文字列のスライスを作成し、初期値として["apple", "banana", "cherry"]を設定する
	slice := []string{"grape", "strawbery", "cherry", "apple", "lemon", "banana"}

	// 2. スライスの要素数が10未満の場合は、スライスに"durian"を追加する
	if len(slice) < 10 {
		slice = append(slice, "durian")
	}

	fmt.Println(slice)
	// 3. スライスの2番目の要素を出力する
	fmt.Println(slice[2])

	// 4. スライスの最後の要素を削除する
	// sliceEnd := len(slice) - 1
	// slice = slice[:sliceEnd+copy(slice[sliceEnd:], slice[sliceEnd+1:])]
	slice = slice[:len(slice)-1]

	// 5. スライスの要素をアルファベット順にソートする
	sort.Strings(slice)
	fmt.Println(slice)
	// 6. スライスに"grape"が含まれているかどうかを確認し、結果を出力する
	for i, name := range slice {
		if name == "grape" {
			fmt.Printf("%d番目に%sが!!!\n", i, name)
		}
	}

	// 7. スライスの要素をカンマ区切りの文字列に変換して出力する
	fmt.Println(strings.Join(slice, ","))
}
