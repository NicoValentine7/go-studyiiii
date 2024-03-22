package main

import "fmt"

func main() {
	// 1. int型の変数xを宣言し、値を10で初期化する
	x := 10

	// 2. 変数xのポインタを保持するポインタ変数ptrを宣言する
	ptr := &x

	// 3. ポインタ変数ptrの値（アドレス）を出力する
	fmt.Println(ptr)
	// 4. ポインタ変数ptrをデリファレンスして、その値を出力する
	fmt.Println(*ptr)
	// 5. ポインタ変数ptrをデリファレンスして、その値を20に変更する
	*ptr = 20
	// 6. 変数xの値を出力する
	fmt.Println(x)
	// 7. 変数xのポインタをダイレクトに取得し、そのポインタをデリファレンスして、値を30に変更する
	// *x = 30 //わからん
	*(&x) = 30
	//　あるいは x = 30

	// 8. 変数xの値を出力する
	fmt.Println(x)
}
