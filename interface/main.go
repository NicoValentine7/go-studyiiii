package main

import "fmt"

// 以下の interface を定義してください
// 1. "Athlete" という名前の interface を定義する
//    - Train() メソッドを持つ
//    - Compete() メソッドを持つ

// ここで interface を定義する
type Athlete interface {
	Train()
	Compete()
}

// 以下の構造体を定義してください
// 1. "Swimmer" という名前の構造体を定義する
//    - name フィールドを持つ（型: string）
//    - style フィールドを持つ（型: string）
// 2. "Runner" という名前の構造体を定義する
//    - name フィールドを持つ（型: string）
//    - distance フィールドを持つ（型: string）

// ここで構造体を定義する
type Swimmer struct {
	name  string
	style string
}

type Runner struct {
	name     string
	distance string
}

// 以下のメソッドを定義してください
// 1. "Swimmer" 構造体に Train() メソッドを実装する
//    - "〇〇 is training in the pool." と出力する
// 2. "Swimmer" 構造体に Compete() メソッドを実装する
//    - "〇〇 is competing in a 〇〇 race." と出力する
// 3. "Runner" 構造体に Train() メソッドを実装する
//    - "〇〇 is training on the track." と出力する
// 4. "Runner" 構造体に Compete() メソッドを実装する
//    - "〇〇 is competing in a 〇〇 race." と出力する

// ここでメソッドを実装する
func (s Swimmer) Train() {
	fmt.Printf("%s is training in the pool.\n", s.name)
}

func (s Swimmer) Compete() {
	fmt.Printf("%s is competing in a %s race.\n", s.name, s.style)
}

func (r Runner) Train() {
	fmt.Printf("%s is training on the track.\n", r.name)
}

func (r Runner) Compete() {
	fmt.Printf("%s is competing in a %s race.\n", r.name, r.distance)
}

func main() {
	// 以下の処理を行ってください
	// 1. Swimmer 構造体と Runner 構造体のインスタンスを作成する
	// 2. 各インスタンスの Train() メソッドと Compete() メソッドを呼び出す
	// 3. 作成したインスタンスを Athlete 型のスライスに格納する
	// 4. スライスの各要素に対して、Train() メソッドと Compete() メソッドを呼び出す

	// ここでインスタンスの作成と呼び出しを行う
	s := Swimmer{
		name:  "Yusaku",
		style: "swim",
	}

	s.Train()
	s.Compete()

	r := Runner{
		name:     "Yusaku２号",
		distance: "400m",
	}

	r.Train()
	r.Compete()
}
