/*package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("my.txt")
	if err != nil {
		log.Fatal("installing my.txt is in your future")
	}
	fmt.Printf("my.txt is available at %s\n", path)
}*/
package main

import (
	"fmt"
)

/*func input1(a *int, b *int) {
	fmt.Println("-------input()----------")
	fmt.Printf("&a:%v, a:%v, *a:%v\n", &a, a, *a)
	fmt.Printf("&b:%v, b:%v, *b:%v\n", &b, b, *b)

	*a = 3
	*b = 4

	fmt.Printf("&a:%v, a:%v, *a:%v\n", &a, a, *a)
	fmt.Printf("&b:%v, b:%v, *b:%v\n", &b, b, *b)
}

func main() {
	fmt.Println("------------main--------------")
	var ia int = 1
	var ib int = 2
	fmt.Printf("&ia:%v, ia:%v\n", &ia, ia)
	fmt.Printf("&ib:%v, ib:%v\n", &ib, ib)

	input(&ia, &ib)

	fmt.Println("-------------main--------------")
	fmt.Printf("&ia:%v, ia:%v\n", &ia, ia)
	fmt.Printf("&ib:%v, ib:%v\n", &ib, ib)
}
*/

/*func modify(a *int) {
	fmt.Println("---------modify------------")
	fmt.Printf("&a:%v, a:%v, *a:%v\n", &a, a, *a)
	a = nil
	fmt.Printf("&a:%v, a:%v\n", &a, a)
}

func main() {
	a := new(int)
	*a = 10
	fmt.Println("------------main-------------")
	fmt.Printf("&a:%v, a:%v, *a:%v\n", &a, a, *a)
	fmt.Println(a)
	modify(a)
	fmt.Println("-------------main-------------")
	fmt.Printf("&a:%v, a:%v, *a:%v\n", &a, a, *a)
}*/

/*type Interface interface {
	say() string
}

type Object struct {
}

func (this *Object) say() string {
	return "hello turtle."
}

func do(i Interface) string {
	return i.say()
}

func main() {
	o := Object{}
	var a float64 = 1.23
	fmt.Println(do(&o))
	fmt.Printf("CCCCCCCCCCCC:%T", a)
}*/

/*var user = ""

func inita() {
	defer func() {
		fmt.Print("defer##\n")
	}()

	if user == "" {
		fmt.Print("@@@before panic\n")
		panic("no value for user\n")
		fmt.Print("!!after panic\n")
	}
}

func throwPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Print(x)
			b = true
		}
	}()

	f()

	fmt.Print("after the func run\n")

	return
}

func main() {
	throwPanic(inita)
	fmt.Println("ok")
}*/

type person struct {
	name string
	Age  int
}

func (p *person) say(in int) {
	p.Age = 2
	fmt.Println("in say", p.Age, in)
}

func (p person) shuo() {
	p.Age = 2
	fmt.Println("who are you ?")
}

func main() {
	// p1 := &person{"turtle", 24}
	p3 := person{"turtle", 24}
	/*	(&p3).say(1)
		fmt.Println("-----------")
		fmt.Println("in main ", p3.Age)
		fmt.Println("-----------")*/
	p3.say(2)
	fmt.Println("-----------")
	fmt.Println("in main ", p3.Age)

	// p2 := person{"turtle", 100}
	/*fmt.Printf("p is %T\n", p)
	p.say()
	fmt.Printf("name is %v, age is %v", p.name, p.age)
	*/
	/*	p1.say(3)
		p2.shuo()
		fmt.Println(p1.Age, p2.Age)*/
}
