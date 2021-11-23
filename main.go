package main

import (
	"fmt"

	"github.com/aiharanaoya/golang-test/foo"
)

func main() {
	fmt.Println("hello world")

	package1()
}

func test() {
	// 変数

	// varで定義する。（型省略も出来る）
	var msg1 string = "hello"
	fmt.Println(msg1)

	// varを使わないパターン
	msg2 := "hello"
	fmt.Println(msg2)

	// データ型
	var msg3 string = "hello"
	var num1 int = 10
	var num2 float64 = 12.34
	var isTrue bool = true
	fmt.Println(msg3, num1, num2, isTrue)

	// 定数
	const msg4 = "hello"
	fmt.Println(msg4)

	fmt.Println(swap(2, 5))

	// 関数を変数に代入するパターン
	f := func(a int, b int) (int, int) {
		return b, a
	}

	fmt.Println(f(3, 8))

	// 即時関数のように宣言して実行
	func(msg string) {
		fmt.Println(msg)
	}("hello")
}

// 複数の戻り値
func swap(a int, b int) (int, int) {
	return b, a
}

// 名前付きreturn関数
func getHelloMessage(name string) (msg string) {
	msg = "Hello " + name
	return
}

func pointer() {
	n := 100

	fmt.Println(n)
	fmt.Println(&n)

	p := &n
	fmt.Println(p)
	fmt.Println(*p)

	*p = 300
	fmt.Println(n)
	fmt.Println(*p)

	n = 200
	fmt.Println(n)
	fmt.Println(*p)

	double(&n)
	fmt.Println(n)
	fmt.Println(*p)

	sl := []int{1, 2, 3}
	double2(sl)
	fmt.Println(sl)
}

func double(i *int) {
	*i = *i * 2
}

func double2(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

type User struct {
	Name string
	Age  int
}

func struct1() {
	var user1 User
	fmt.Println(user1)

	user1.Name = "user1"
	user1.Age = 20
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)

	user2.Name = "user2"
	user2.Age = 25
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	// newをつけるとポインタ型になる
	user4 := new(User)
	fmt.Println(user4)

	// ポインタ型になる、こっちの方が使われる
	user5 := &User{}
	fmt.Println(user5)

	// ポインタ型で渡すか、実体を渡すか
	// アドレスを渡すと参照渡し
	// 実体を渡すと値渡し
	userA := User{}
	userB := &User{}

	updateUser(userA)
	updateUser2(userB)

	fmt.Println(userA)
	fmt.Println(userB)
}

func updateUser(user User) {
	user.Name = "updateUser"
	user.Age = 100
}

func updateUser2(user *User) {
	user.Name = "updateUser"
	user.Age = 100
}

type User2 struct {
	Name string
	Age  int
}

type Product struct {
	id    int
	value int
}

type T struct {
	User2
	Product
}

type Products []*Product

func (u User2) SayName() {
	fmt.Println(u.Name)
}

// レシーバーはポインタ型にしておくのが望ましい
func (u *User2) SetName(name string) {
	u.Name = name
}

func NewUser(name string, age int) *User2 {
	return &User2{Name: name, Age: age}
}

func struct2() {
	user1 := User2{Name: "ユーザー"}
	user1.SayName()

	user1.SetName("セットユーザー")
	user1.SayName()

	t := T{
		User2:   User2{Name: "ああ", Age: 20},
		Product: Product{id: 1, value: 300},
	}

	fmt.Println(t)
	fmt.Println(t.User2)
	fmt.Println(t.Product)

	t.User2.SetName("Tユーザー")
	t.User2.SayName()

	user2 := NewUser("ユーザー2", 25)
	fmt.Println(*user2)

	product1 := Product{id: 1, value: 1000}
	product2 := Product{id: 2, value: 1200}
	product3 := Product{id: 3, value: 1400}

	products := Products{}

	products = append(products, &product1, &product2, &product3)

	fmt.Println(products)

	for _, p := range products {
		fmt.Println(*p)
	}

	m := map[int]Product{
		1: {id: 1, value: 1000},
		2: {id: 2, value: 1200},
	}

	fmt.Println(m)

	m2 := map[Product]string{
		{id: 1, value: 1000}: "商品1",
		{id: 2, value: 1200}: "商品2",
	}

	fmt.Println(m2)
}

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラー発生", ErrCode: 1}
}

func interface1() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "1234", Model: "AAAA"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	err := RaiseError()
	fmt.Println(err.Error())
}

func package1() {
	fmt.Println(foo.Max)
}
