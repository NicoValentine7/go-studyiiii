Go言語にinterfaceが導入された背景には、言語設計者たちがシンプルで柔軟性の高いオブジェクト指向プログラミングの仕組みを提供したいという思想がありました。当時、他の言語ではクラスベースの継承が主流でしたが、Go言語ではより軽量で理解しやすい方法を模索していました。そこで、interfaceによるダック・タイピング（Duck Typing）の概念を取り入れることにしたのです。

ダック・タイピングとは、「もしもそれがアヒルのように歩き、アヒルのように鳴くのなら、それはアヒルであろう」という考え方です。つまり、オブジェクトの型よりも、そのオブジェクトが持つ振る舞い（メソッド）に着目するということです。

例えば、動物園を管理するプログラムを作るとします。このプログラムでは、様々な動物（ゾウ、キリン、ライオンなど）を扱う必要があります。これらの動物たちは、それぞれ固有の特徴を持っていますが、共通の行動（鳴く、食べる、寝るなど）もあります。

従来のクラスベースの言語では、これらの動物たちを表現するために、「動物」という基底クラスを作り、そこから各動物クラスを継承させるという方法を取ることが多いでしょう。しかし、Go言語のinterfaceを使えば、以下のようにシンプルに表現できます：

```go
type Animal interface {
    Speak() string
    Eat(food string)
    Sleep()
}
```

このAnimal interfaceは、Speak、Eat、Sleepという3つのメソッドを定義しています。そして、以下のようにElephant、Giraffe、Lionという3つの構造体を定義し、それぞれがAnimal interfaceを実装します：

```go
type Elephant struct{}

func (e Elephant) Speak() string {
    return "Pawoo!"
}

func (e Elephant) Eat(food string) {
    fmt.Println("Elephant eats", food)
}

func (e Elephant) Sleep() {
    fmt.Println("Elephant sleeps")
}

type Giraffe struct{}

func (g Giraffe) Speak() string {
    return "..."
}

func (g Giraffe) Eat(food string) {
    fmt.Println("Giraffe eats", food)
}

func (g Giraffe) Sleep() {
    fmt.Println("Giraffe sleeps")
}

type Lion struct{}

func (l Lion) Speak() string {
    return "Roar!"
}

func (l Lion) Eat(food string) {
    fmt.Println("Lion eats", food)
}

func (l Lion) Sleep() {
    fmt.Println("Lion sleeps")
}
```

これで、ElephantもGiraffeもLionもみなAnimal interfaceを実装しているので、これらのインスタンスをAnimal型の変数に代入することができます。

```go
var a Animal

a = Elephant{}
fmt.Println(a.Speak())
a.Eat("grass")
a.Sleep()

a = Giraffe{}
fmt.Println(a.Speak())
a.Eat("leaves")
a.Sleep()

a = Lion{}
fmt.Println(a.Speak())
a.Eat("meat")
a.Sleep()
```

このように、Go言語のinterfaceを使えば、異なる型のオブジェクトを同じように扱うことができます。これにより、コードの柔軟性と拡張性が高まり、より理解しやすいオブジェクト指向プログラミングが可能になるのです。
