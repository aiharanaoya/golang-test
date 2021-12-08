package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/aiharanaoya/golang-test/alib"
	"github.com/aiharanaoya/golang-test/foo"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("hello world")

	dbTest1()
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
	fmt.Println(foo.ReturnMin())
}

// テスト
func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}

	// return i == 1
}

func alib1() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
}

func study1() {
	// if文
	a := 2
	if a == 1 {
		fmt.Println("one")
	} else if a == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("other")
	}

	// if a := 2; a == 1 {
	// 	fmt.Println("one")
	// } else if a == 2 {
	// 	fmt.Println("two")
	// } else {
	// 	fmt.Println("other")
	// }

	// switch文
	switch b := 2; b {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}

	switch score := 66; {
	case score > 80:
		fmt.Println("Great!")
	case score > 60:
		fmt.Println("Good!")
	default:
		fmt.Println("Bad")
	}

	// for文
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		} else if i == 8 {
			break
		}
		fmt.Println(i)
	}

	c := 0
	for c < 10 {
		fmt.Println(c)
		c++
	}

	// 配列
	var ar1 [5]int
	ar1[1] = 3
	ar1[2] = 4
	fmt.Println(ar1)

	ar2 := [5]int{2, 3, 4, 5, 6}
	fmt.Println(ar2)

	ar3 := [...]int{2, 3, 4, 5, 6}
	fmt.Println(ar3)

	// スライス
	ar := [5]int{2, 3, 4, 5, 6}

	sl1 := ar[2:4]
	fmt.Println(sl1)

	sl2 := ar[2:]
	fmt.Println(sl2)

	sl3 := ar[:4]
	fmt.Println(sl3)

	sl4 := ar[:]
	fmt.Println(sl4)

	sl5 := []int{2, 3, 4, 5, 6}
	fmt.Println(sl5)

	sl6 := make([]int, 5)
	fmt.Println(sl6)

	sl7 := append(sl5, 7, 8)
	fmt.Println(sl7)

	fmt.Println(len(sl7))

	sl8 := sl7
	fmt.Println(sl8)

	sl8[0] = 100
	fmt.Println(sl7, sl8)

	sl9 := make([]int, 7)
	copy(sl9, sl8)
	fmt.Println(sl9)

	sl9[0] = 200
	fmt.Println(sl8, sl9)

	// マップ
	var ma1 = map[string]int{"A": 100, "B": 200}
	fmt.Println(ma1)

	ma1["A"] = 1000
	ma1["B"] = 2000
	fmt.Println(ma1)

	ma2 := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	fmt.Println(ma2)

	rangeArr := [3]int{1, 2, 3}

	for i, v := range rangeArr {
		fmt.Println(i, v)
	}

	rangeSl := []int{1, 2, 3}
	for _, v := range rangeSl {
		fmt.Println(v)
	}

	// ポインタ
	var nu1 int = 100
	fmt.Println(nu1)

	fmt.Println(&nu1)

	var po1 *int = &nu1
	fmt.Println(po1)

	fmt.Println(*po1)

	double3(nu1)
	fmt.Println(nu1)

	double4(&nu1)
	fmt.Println(nu1)

	arPo := [3]int{1, 2, 3}
	arrayPo(arPo)
	fmt.Println(arPo)

	slPo := []int{1, 2, 3}
	slicePo(slPo)
	fmt.Println(slPo)

	test1()
}

func double3(i int) {
	i = i * 2
}

func double4(i *int) {
	*i = *i * 2
}

func arrayPo(i [3]int) {
	i[0] = 100
}

func slicePo(i []int) {
	i[0] = 100
}

func test1() {
	sum := 0
	square := 0

	for i := 1; i <= 100; i++ {
		sum += (i * i)
		square += i
	}

	square = (square * square)

	ans := square - sum

	fmt.Println(ans)
}

func lib() {
	f1, err := os.Open("test.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer f1.Close()

	fmt.Println(os.Args[1])

	// f2, _ := os.Create("foo.txt")

	// f2.Write([]byte("Hello\n"))
	// f2.WriteAt([]byte("Golang"), 6)
	// f2.Seek(0, os.SEEK_END)
	// f2.WriteString("AAA\nBBB")

	// 今の時間
	t1 := time.Now()
	fmt.Println(t1)

	// 指定時間を生成
	t2 := time.Date(2021, 11, 29, 20, 0, 0, 0, time.Local)
	fmt.Println(t2)
	fmt.Println(t2.Year())
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday())
	fmt.Println(t2.Day())

	// t1 - t2
	d1 := t1.Sub(t2)
	fmt.Println(d1)

	// 円周率
	fmt.Println(math.Pi)

	// 絶対値
	fmt.Println(math.Abs(-100))

	// 累乗
	fmt.Println(math.Pow(2, 3))

	// 平方根
	fmt.Println(math.Sqrt(2))
}

// type MyHandler struct{}

// func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World")
// }

// net/http server
func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Panicln(err)
	}
	t.Execute(w, "Hello World")
}

func server() {
	http.HandleFunc("/top", top)
	http.ListenAndServe(":8080", nil)
}

/////////////////////////////////////

// DB sqlite3
var Db *sql.DB

func dbTest1() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	// テーブルの作成
	// cmd := `CREATE TABLE IF NOT EXISTS persons (
	// 					name STRING,
	// 					age INT)`
	//
	// _, err := Db.Exec(cmd)

	// データの追加
	// cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	//
	// _, err := Db.Exec(cmd, "Toru", 25)

	// データの更新
	// cmd := "UPDATE persons SET age = ? WHERE name = ?"

	// _, err := Db.Exec(cmd, 30, "Taro")

	// データの削除
	cmd := "DELETE FROM persons WHERE name = ?"

	_, err := Db.Exec(cmd, "Toru")

	if err != nil {
		log.Fatalln(err)
	}

	// データの取得
	// cmd := "SELECT * FROM persons WHERE age = ?"

	// QueryRow: 1レコード取得
	// row := Db.QueryRow(cmd, 25)

	// var p Person

	// err := row.Scan(&p.Name, &p.Age)

	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No Row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }

	// fmt.Println(p.Name, p.Age)
}

////////////////////////////////////////
