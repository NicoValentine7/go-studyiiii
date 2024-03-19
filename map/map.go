package main

import "fmt"

func main() {
	// 1. string型のキーとint型の値を持つマップを作成する
	fruits := make(map[string]int)

	// 2. マップに以下のキーと値のペアを追加する
	//    "apple" → 150
	//    "banana" → 200
	//    "cherry" → 300
	fruits["apple"] = 150
	fruits["banana"] = 200
	fruits["cherry"] = 300

	fmt.Println(fruits)
	// 3. マップから"banana"の値を取得して出力する
	fmt.Printf("banana map is: %d\n", fruits["banana"])

	// 4. マップに"durian"というキーが存在するかを確認し、存在する場合はその値を出力する
	//    存在しない場合は"not found"と出力する
	if price, ok := fruits["durian"]; ok {
		fmt.Println("Price is Durian:", price)
	} else {
		fmt.Println("not found")
	}

	// 5. マップの要素数を出力する
	fmt.Println(len(fruits))

	// 6. マップから"apple"を削除する
	delete(fruits, "apple")

	// 7. マップのすべてのキーと値を出力する
	fmt.Printf("fruits: %v", fruits)
}
