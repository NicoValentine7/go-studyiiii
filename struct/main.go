package main

import "fmt"

//  1. "Person"という名前の構造体を定義する
//     フィールドは以下の通り
//     - Name: string
//     - Age: int
//     - Address: string
type Person struct {
	Name    string
	Age     int
	Address string
}

//  2. "NewPerson"というコンストラクタ関数を定義する
//     引数でName, Age, Addressを受け取り、Personのインスタンスを返す
func NewPerson(name string, age int, address string) *Person {
	return &Person{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

func main() {
	// 3. NewPersonを使って、Personのインスタンスを作成する
	//    名前: "Alice", 年齢: 25, 住所: "123 Main St"
	person := NewPerson("Alice", 25, "123 Main St")

	// 4. 作成したインスタンスのNameとAgeを出力する
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)

	// 5. 作成したインスタンスのAddressを"456 Elm St"に更新する
	person.Address = "456 Elm St"

	// 6. 更新後のインスタンスの情報を出力する
	fmt.Println(person)
}
