package main

import "fmt"

//  1. "Rectangle"という名前の構造体を定義する
//     フィールドは以下の通り
//     - Width: float64
//     - Height: float64
type Rectangle struct {
	Width  float64
	Height float64
}

//  2. "NewRectangle"というコンストラクタ関数を定義する
//     引数でWidthとHeightを受け取り、Rectangleのインスタンスを返す
func NewRectangle(width float64, height float64) *Rectangle {
	return &Rectangle{
		Width:  width,
		Height: height,
	}
}

//  3. "Rectangle"構造体に"Area"というメソッドを定義する
//     メソッドは長方形の面積を計算して返す

// 不正解
// func (*Rectangle) Area() float64 {
// 	rectangle := NewRectangle(30, 40)
// 	return rectangle.Height * rectangle.Width
// }

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

//  4. "Rectangle"構造体に"Perimeter"というメソッドを定義する
//     メソッドは長方形の周囲の長さを計算して返す
func (r *Rectangle) Perimeter() float64 {
	return r.Height*2 + r.Width*2
}

func main() {
	// 5. NewRectangleを使って、Rectangleのインスタンスを作成する
	//    幅: 5.0, 高さ: 3.0
	r := NewRectangle(5.0, 3.0)

	// 6. 作成したインスタンスのWidthとHeightを出力する
	fmt.Println("Width = ", r.Width)
	fmt.Println("Height = ", r.Height)

	// 7. 作成したインスタンスのAreaメソッドを呼び出して、面積を出力する
	fmt.Println("Area = ", r.Area())

	// 8. 作成したインスタンスのPerimeterメソッドを呼び出して、周囲の長さを出力する
	fmt.Println("Perimeter = ", r.Perimeter())
}
