package main

import (
	"context"
	"database/sql"
	json2 "encoding/json"
	"flag"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"udemy/alib"
	"udemy/foo"
)

//i5 := 100
var i5 int = 500

const Pi = 3.14

//const (
//	URL = ""
//)

const (
	c1 = iota
	c2
	c3
)

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	// fmt.Println("Hello World")
	// fmt.Println(time.Now())

	//var i int = 100
	//fmt.Println(i)
	//
	//var s string = "Hello Go"
	//fmt.Println(s)
	//
	//var t, f bool = true, false
	//fmt.Println(t, f)
	//
	//var (
	//	i2 int    = 200
	//	s2 string = "Golang"
	//)
	//fmt.Println(i2, s2)
	//
	//var i3 int
	//var s3 string
	//fmt.Println(i3, s3)
	//
	//i3 = 300
	//s3 = "Go"
	//fmt.Println(i3, s3)
	//
	//i = 150
	//fmt.Println(i)
	//
	//i4 := 400
	//fmt.Println(i4)
	//
	//i4 = 450
	//fmt.Println(i4)
	//
	//// i4 = "Hello"
	//// fmt.Println(i4)
	//
	//fmt.Println(i5)
	//
	//outer()
	//
	//var s5 string = "Not Use"
	//fmt.Println(s5)

	var i int = 100
	var i2 int64 = 200

	fmt.Println(i + 50)

	//fmt.Println(i + i2)

	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))

	var f float64 = 2.4
	fmt.Println(f)

	fl := 3.2
	fmt.Println(f + fl)
	fmt.Printf("%T, %T\n", f, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	c := []byte("Hi")
	fmt.Println(c)

	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr1[0])

	//var arr5 = []string{"1","2"}
	//var arr6 [4]int

	arr2[2] = "c"
	fmt.Println(len(arr2))

	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = 1.2
	fmt.Println(x)
	x = "aa"
	fmt.Println(x)
	x = [3]int{2, 3}
	fmt.Println(x)

	//x = 2
	//fmt.Println(x + 3)

	var i4 int = 1
	fl64 := float64(i4)
	fmt.Println(fl64)
	fmt.Printf("fl64 = %T\n", fl64)

	i3 := int32(fl64)
	fmt.Printf("i3 = %T\n", i3)

	var s string = "100"
	i5, _ := strconv.Atoi(s)
	fmt.Println(i5)
	fmt.Printf("i5 = %T\n", i5)
	var i6 int = 200
	s2 := strconv.Itoa(i6)
	fmt.Printf("s2 = %T\n", s2)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	var h2 string = string(b)
	fmt.Println(h2)

	fmt.Println(Pi)
	fmt.Println(c1, c2, c3)

	i7 := Plus(1, 2)
	fmt.Println(i7)

	i8, _ := Div(9, 4)
	fmt.Println(i8)

	i10 := Double(1000)
	fmt.Println(i10)

	Noreturn()

	fn := func(x, y int) int {
		return x + y
	}
	i11 := fn(1, 2)
	fmt.Println(i11)

	i12 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(i12)

	f2 := ReturnFunc()
	f2()

	CallFunc(func() {
		fmt.Println("Call Func")
	})

	f3 := Later()
	fmt.Println(f3("Hello"))
	fmt.Println(f3("My"))
	fmt.Println(f3("name"))
	fmt.Println(f3("is"))
	fmt.Println(f3("Golang"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())

	i = 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("Loop")
	}

	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	for i11 := 0; i11 < 10; i11++ {
		if i11 == 3 {
			continue
		}
		if i11 == 6 {
			break
		}
		fmt.Println(i11)
	}

	arr5 := [3]int{1, 2, 3}
	for i = 0; i < len(arr5); i++ {
		fmt.Println(arr5[i])
	}

	for i, v := range arr5 {
		fmt.Println(i, v)
	}

	sl := []string{"Python", "PHP", "GO"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for i, v := range m {
		fmt.Println(i, v)
	}

	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("i dont konw")
	}

	switch {
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4")
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	}

	anything("aaa")
	anything(1)

	var xx interface{} = 3
	i13 := xx.(int)
	fmt.Println(i13 + 2)

	ff, isFloat64 := xx.(float64)
	fmt.Println(ff, isFloat64)

	if xx == nil {
		fmt.Println("none")
	} else if i, isInt := xx.(int); isInt {
		fmt.Println(i, "xx is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "xx is string")
	}

	switch vv := xx.(type) {
	case int:
		fmt.Println(vv, "int")
	case string:
		fmt.Println(vv, "string")
	}

Loop:
	for {
		for {
			for {
				fmt.Println("Start")
				break Loop
			}
			fmt.Println("no operation")
		}
		fmt.Println("no operation")
	}
	fmt.Println("End")
	//TestDefer()
	//
	//defer func() {
	//	fmt.Println("1")
	//	fmt.Println("2")
	//	fmt.Println("3")
	//}()
	//
	//RunDefer()
	//
	//file, err := os.Create("test.txt")
	//if err != nil {
	//	fmt.Println("ERROR")
	//}
	//defer file.Close()
	//
	//file.Write([]byte("Hello"))

	//defer func() {
	//	if x:= recover(); x != nil {
	//		fmt.Println(x)
	//	}
	//}()
	//panic("runtime error")
	//fmt.Println("Start")

	//go sub()
	//go sub()
	//
	//for {
	//	fmt.Println("Main Loop")
	//	time.Sleep(200 * time.Microsecond)
	//}

	fmt.Println("Main")

	var sl1 []int
	fmt.Println(sl1)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	var sl3 []string = []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[0])

	fmt.Println(sl5[2:4])

	fmt.Println(sl5[:2])
	fmt.Println(sl5[2:])

	fmt.Println(sl5[:])

	fmt.Println(sl5[len(sl5)-1])

	fmt.Println(sl5[1 : len(sl5)-1])

	sl6 := []int{100, 200}
	fmt.Println(sl6)

	sl6 = append(sl6, 300)
	fmt.Println(sl6)

	sl6 = append(sl6, 400, 500, 600)
	fmt.Println(sl6)

	sl7 := make([]int, 5)
	fmt.Println(sl7)
	fmt.Println(len(sl7))
	fmt.Println(cap(sl7))

	sl8 := make([]int, 5, 10)
	fmt.Println(len(sl8))
	fmt.Println(cap(sl8))

	sl8 = append(sl8, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl8))
	fmt.Println(cap(sl8))

	sl9 := []int{100, 200}
	sl10 := sl9

	sl10[0] = 1000
	fmt.Println(sl9)

	sl11 := []int{1, 2, 3, 4, 5}
	sl12 := make([]int, 5, 10)
	nn := copy(sl12, sl11)
	fmt.Println(nn, sl12)

	sl13 := []string{"A", "B", "C"}
	fmt.Println(sl13)

	for i, v := range sl13 {
		fmt.Println(i, v)
	}

	for i := 0; i < len(sl13); i++ {
		fmt.Println(sl13[i])
	}

	fmt.Println(sum(1, 2, 3))

	sl14 := []int{1, 2, 3}
	fmt.Println(sum(sl14...))

	var m1 = map[string]int{"A": 100, "B": 200}
	fmt.Println(m1)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	ss, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(ss, ok)

	delete(m4, 2)
	fmt.Println(m4)

	m5 := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}
	for k, v := range m5 {
		fmt.Println(k, v)
	}

	var ch1 chan int

	// receive only
	//var ch2 <-chan int

	// send only
	//var ch3 chan<- int

	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 1
	fmt.Println(len(ch3))

	ch3 <- 2
	fmt.Println(len(ch3))

	i14 := <-ch3
	fmt.Println(i14)
	i14 = <-ch3
	fmt.Println(i14)
	fmt.Println(len(ch3))

	//ch4 := make(chan int)
	//ch5 := make(chan int)

	//go receiver(ch4)
	//go receiver(ch5)

	//i15 := 0
	//for i15 < 100 {
	//	ch4 <- i15
	//	ch5 <- i15
	//	time.Sleep(50 * time.Millisecond)
	//	i15++
	//}

	//ch6 := make(chan int, 2)

	//close(ch6)
	//ch6 <- 1
	//fmt.Println(<-ch6)
	//i16, ok := <- ch6
	//fmt.Println(i16, ok)
	//fmt.Println(<-ch4)

	ch17 := make(chan int, 2)
	go receiver2("1.goroutine", ch17)
	go receiver2("2.goroutine", ch17)

	i18 := 0
	for i18 < 100 {
		ch17 <- i18
		i18++
	}
	close(ch17)
	time.Sleep(3 * time.Second)

	ch18 := make(chan int, 3)
	ch18 <- 1
	ch18 <- 2
	ch18 <- 3
	close(ch18)
	for i := range ch18 {
		fmt.Println(i)
	}

	ch19 := make(chan int, 2)
	ch20 := make(chan string, 2)

	/*
		ch19 <- 1
		ch20 <- "A"
	*/
	/*
		v1 := <-ch19
		v2 := <-ch20
	*/

	select {
	case v1 := <-ch19:
		fmt.Println(v1 + 1000)
	case v2 := <-ch20:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("no select")
	}

	ch21 := make(chan int)
	ch22 := make(chan int)
	ch23 := make(chan int)

	go func() {
		for {
			i := <-ch21
			ch22 <- i * 2
		}
	}()

	go func() {
		for {
			i := <-ch22
			ch23 <- i - 1
		}
	}()

	n = 0
	for {
		select {
		case ch21 <- n:
			n++
		case i19 := <-ch23:
			fmt.Println("received", i19)
		}
		if n > 100 {
			break
		}
	}

	var n1 int = 100
	fmt.Println(n1)
	fmt.Println(&n1)

	double(n1)
	fmt.Println(n1)

	var p *int = &n1
	fmt.Println(p)
	fmt.Println(*p)

	/*
		*p = 300
		fmt.Println(n1)
	*/

	double2(&n1)
	fmt.Println(n1)

	double2(p)
	fmt.Println(*p)

	var pp Point
	fmt.Println(pp)

	p2 := Point{A: 1, B: "GO", C: 1.2}
	fmt.Println(p2)
	fmt.Println(p2.A)

	p3 := Point{1, "Python", 3.14}
	fmt.Println(p3)

	p4 := Point{}
	update(&p4)
	fmt.Println(p4)

	p5 := &Point{}
	update(p5)
	fmt.Println(*p5)

	p6 := new(Point)
	update(p6)
	fmt.Println(*p6)

	p6.set(3)
	fmt.Println(p6)

	bp := BigPoint{}
	fmt.Println(bp)

	bp.Point.A = 100
	bp.Point.B = "Apple"
	bp.Point.C = 2.8
	fmt.Println(bp)

	p7 := NewPoint(1, "B", 2.2)
	fmt.Println(p7)

	ps := make([]Person, 5)
	fmt.Println(ps)
	ps[0].Name = "Mike"
	ps[1].Name = "John"
	ps[2].Name = "Nina"
	ps[3].Name = "Joe"
	ps[4].Name = "Nancy"
	fmt.Println(ps)

	pss := Persons{}
	ps1 := Person{"Mike"}
	ps2 := Person{"John"}
	ps3 := Person{"Nina"}
	ps4 := Person{"Joe"}
	ps5 := Person{"Nancy"}
	pss.Persons = append(pss.Persons, &ps1, &ps2, &ps3, &ps4, &ps5)

	for _, p := range pss.Persons {
		fmt.Println(p)
	}

	vs := []Stringfy{
		&Person2{"Taro", 21},
		&Car{"123-456", "AB-1234"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	err := RaiseError()
	fmt.Println(err.Error())

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}

	pp1 := &Point2{100, "ABC"}
	fmt.Println(pp1)

	fmt.Println(foo.Max)

	fmt.Println(Do("AAA"))

	sl15 := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(sl15))

	// deferも実行されない
	//os.Exit(1)
	//fmt.Println("end")

	file, err2 := os.Open("test.txt")
	if err2 != nil {
		log.Fatalln(err2)
	}
	defer file.Close()

	file2, _ := os.Create("foo.txt")
	file2.Write([]byte("Hello\n"))
	file2.WriteAt([]byte("Golang"), 6)

	file2.Seek(0, 2)
	file2.WriteString("Year")
	file2.Close()

	file3, _ := os.OpenFile("foo.txt", os.O_RDONLY, 0666)
	defer file3.Close()

	bs := make([]byte, 128)
	nnn, _ := file3.Read(bs)
	fmt.Println(nnn)
	fmt.Println(string(bs))

	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Year())
	t2 := t.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t2)
	d2 := t.Sub(t2)
	fmt.Println(d2)
	fmt.Println(t.After(t2))

	fmt.Println(math.Pi)

	var (
		max int
		msg string
		flg bool
	)

	flag.IntVar(&max, "n", 32, "max value")
	flag.StringVar(&msg, "m", "", "message")
	flag.BoolVar(&flg, "x", false, "flag")

	flag.Parse()

	fmt.Println(max, msg, flg)

	log.SetOutput(os.Stdout)
	log.Println("Log")
	log.Printf("%d", 3)

	//log.Fatal("Log")

	//log.Panic("Log")

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("Log")

	log.SetPrefix("[LOG]")
	log.Println("Log")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("message")

	ss1 := strings.Join([]string{"A", "B", "C"}, ",")
	fmt.Println(ss1)

	ii1 := strings.Index("ABCDE", "A")
	ii2 := strings.Index("ABCDE", "F")
	fmt.Println(ii1, ii2)

	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}

	file1, _ := os.Open("foo.txt")
	bs1, _ := ioutil.ReadAll(file1)
	fmt.Println(string(bs1))

	mutex = new(sync.Mutex)
	wg := new(sync.WaitGroup)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				UpdateAndPrint(i)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	user := new(User)
	user.ID = 1
	user.Name = "aruga"
	user.Email = "example@domain.jp"
	user.Created = time.Now()

	bs2, _ := json2.Marshal(user)
	fmt.Println(string(bs2))

	fmt.Printf("%T\n", bs2)

	user2 := new(User)
	json2.Unmarshal(bs2, &user2)
	fmt.Println(user2)

	isort := []int{5, 2, 3, 2, 3, 4, 5, 3, 1, 2, 3, 4}
	sort.Ints(isort)
	fmt.Println(isort)

	channel := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go loadProcess(ctx, channel)

L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Error")
			fmt.Println(ctx.Err())
			break L
		case s := <-channel:
			fmt.Println(s)
			fmt.Println("Success")
			break L
		}
	}
	fmt.Println("Loop Exit")

	u, _ := url.Parse("http://example.com/search?a=123&b=456#top")
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)
	fmt.Println(u.Query())

	urll := &url.URL{}
	urll.Scheme = "https"
	urll.Host = "google.com"
	q := urll.Query()
	q.Set("q", "Golang")
	urll.RawQuery = q.Encode()
	fmt.Println(urll)

	res, _ := http.Get("https://example.com")
	fmt.Println(res.Status)
	fmt.Println(res.Proto)
	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Content-Type"])
	fmt.Println(res.Request.Method)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	vss := url.Values{}
	vss.Add("id", "1")
	vss.Add("message", "メッセージ")
	fmt.Println(vss.Encode())
	res, _ = http.PostForm("https://example.com", vss)
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	//http.ListenAndServe(":8080", &MyHandler{})
	//http.HandleFunc("/top", top)
	//http.ListenAndServe(":8080", nil)

	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	cmd := `CREATE TABLE IF NOT EXISTS persons(
				name STRING,
				age INT
            )`
	_, err1 := Db.Exec(cmd)
	if err1 != nil {
		log.Fatalln(err1)
	}

	cmd = "INSERT INTO persons (name, age) VALUES (?,?)"
	_, err1 = Db.Exec(cmd, "taro", 20)
	if err1 != nil {
		log.Fatalln(err)
	}

	cmd = "INSERT INTO persons (name, age) VALUES (?,?)"
	_, err1 = Db.Exec(cmd, "hanako", 19)
	if err1 != nil {
		log.Fatalln(err)
	}

	cmd = "SELECT * FROM persons WHERE age = ?"
	row := Db.QueryRow(cmd, 20)
	//fmt.Println(row)
	var person Person1
	err1 = row.Scan(&person.Name, &person.Age)
	if err1 != nil {
		if err1 == sql.ErrNoRows {

		} else {

		}
	}
	fmt.Println(person.Name, person.Age)

	uuidObj, _ := uuid.NewUUID()
	fmt.Println(" ", uuidObj.String())
	uuidObj, _ = uuid.NewRandom()
	fmt.Println(" ", uuidObj.String())
}

type Person1 struct {
	Name string
	Age  int
}

var Db *sql.DB

func top(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "Hello World!!!")
}

func loadProcess(ctx context.Context, ch chan string) {
	fmt.Println("start")
	time.Sleep(2 * time.Second)
	fmt.Println("end")
	ch <- "実行終了"
}

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

type A struct{}
type User struct {
	ID      int `json:"id,string"`
	Name    string
	Email   string
	Created time.Time
	A       A
}

func (u User) MarshalJSON() ([]byte, error) {
	v, error := json2.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, error
}

var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	mutex.Lock()
	st1.A = n
	time.Sleep(time.Microsecond)
	st1.B = n
	time.Sleep(time.Microsecond)
	st1.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st1)
	mutex.Unlock()
}

var st1 struct {
	A, B, C int
}

type Point2 struct {
	A int
	B string
}

func (p *Point2) String() string {
	return fmt.Sprintf("<<%v %v>>", p.A, p.B)
}

func (p *Point) set(i int) {
	p.A = i
}

func update(p *Point) {
	p.A = 100
	p.B = "update"
	p.C = 2.14
}

func RaiseError() error {
	return &MyError{Message: "Custom Error", ErrCode: 1234}
}

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

//type error interface {
//	Error() string
//}

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
}

type Person2 struct {
	Name string
	Age  int
}

type Car struct {
	Number string
	Model  string
}

func (p *Person2) ToString() string {
	return fmt.Sprintf("Name = %v Age = %v", p.Name, p.Age)
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number = %v Model = %v", c.Number, c.Model)
}

type Persons struct {
	Persons []*Person
}

type Point struct {
	A int
	B string
	C float64
}

type BigPoint struct {
	Point
	A int
}

func NewPoint(a int, b string, c float64) *Point {
	return &Point{a, b, c}
}

func double(i int) {
	i *= 2
}

func double2(i *int) {
	*i *= 2
}

func receiver2(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")

}

func receiver(c chan int) {
	for {
		//fmt.Println("Start")
		i := <-c
		fmt.Println(i)
	}
}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func init() {
	fmt.Println("Init")
}

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Microsecond)
	}
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func anything(a interface{}) {
	fmt.Println(a)
}

func Plus(x int, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
}

func ReturnFunc() func() {
	return func() {
		fmt.Println("function!")
	}
}

func CallFunc(f func()) {
	f()
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
