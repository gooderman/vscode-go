package main

import "fmt"
import "math"

/*
first go progaram
*/
// func testLook()
func main() {
	fmt.Println("Hello, World!")
	var a float64
	fmt.Println(a)
	a = math.Abs(-100)
	fmt.Println(a)

	var b string
	b = "string我们🌹"
	fmt.Println(b)

	var c bool
	c = true
	fmt.Println(c)

	aa, bb, cc := 100, "abc", true
	fmt.Printf("%d,%s,%t", aa, bb, cc)
	fmt.Println()
	var aaa, bbb, ccc string
	aaa, bbb, ccc = "123", "aaa", "ddd"

	fmt.Printf("%s,%s,%s", aaa, bbb, ccc)

	testLoop()
	testConst()
	a0, b0, c0 := testFunc(100, "abc")
	fmt.Println(a0, b0, c0)
	testStruct()
}

const (
	CMD_A = iota
	CMD_B
	CMD_C
	CMD_D = 1 << iota
	CMD_E
	CMD_F
)

func testConst() {
	const MAX int32 = 1000
	fmt.Println("const MAX = ", MAX)
	fmt.Println("const CMD_A = ", CMD_A) //0
	fmt.Println("const CMD_B = ", CMD_B) //1
	fmt.Println("const CMD_D = ", CMD_D) //8
	fmt.Println("const CMD_E = ", CMD_E) //16
	fmt.Println("const CMD_F = ", CMD_F) //32

}
func testLoop() {

	var tt []int32 = []int32{1, 2, 3, 4, 5}
	fmt.Printf("tt len = %d\n", len(tt))
	for i, t := range tt {
		fmt.Println(i, t)
	}
	i := 0
	for {
		i++
		if i > 10 {
			break
		}
		fmt.Print(i, "-")
	}
	fmt.Println()
}

type node struct {
	data int
}

func (p *node) add(a int) int {
	p.data += a
	return p.data
}
func testStruct() {
	n := node{100}
	m := &n
	m.add(200)
	n.add(300)
	fmt.Println(m.data)

	pn := new(node)
	pn.data = 200
	pn.add(50)
	fmt.Println(pn.data)

	ppn := &node{3333}
	fmt.Println(ppn.data)

}
func testFunc(a int, b string) (aa int, bb string, cc string) {

	var tt []int32 = []int32{1, 2, 3, 4, 5}
	fmt.Printf("tt len = %d\n", len(tt))
	for i, t := range tt {
		fmt.Println(i, t)
	}
	return a, b, b + b
}
