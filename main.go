package main

import "fmt"

var x = "Hello world!"

func main() {
	x = "World"
	fmt.Println(x)
	forfunc()
	massive1()
	//slice()
	testfunc()
	minimum()
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	defer second() //5
	defer first()  //4
	defer third()  //3
	defer fourth() //2
	defer fifth()  //1
	//54312

	xm := make([]int, 5)
	for i := 0; i < len(xm); i++ {
		fmt.Print("Enter the : ", i, " element: ")
		fmt.Scanf("%d", &xm[i])

	}
	fmt.Println(sum(xm...))
}

func f() {
	fmt.Print("Enter a fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println("Result = ", (input-32)*5/9)
}

func forfunc() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
	i := 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}

func massive() {
	var x [2]string
	x[1] = "sda"
	x[0] = "wwad"
	fmt.Println(x)
}

func massive1() {
	x := [5]float64{55, 23, 213, 123, 23}
	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

func slice() {
	fmt.Print("Enter the size: ")
	var size int
	fmt.Scanf("%d", &size)
	fmt.Println(size)
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	slice3 := make([]int, size)
	fmt.Println(len(slice3))
	for i := 0; i < len(slice3); i++ {
		fmt.Print("Enter the : ", i, " element: ")
		fmt.Scanf("%d", &slice3[i])

	}
	copy(slice2, slice1)
	fmt.Println(slice1, slice2, slice3)
}

func testfunc() {
	fmt.Println(len(make([]int, 3, 9)))
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])
}

func minimum() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := 100
	for i := 0; i < len(x); i++ {
		if min > x[i] {
			min = x[i]
		}
	}
	fmt.Println((min))
}

func average(xs []float64) float64 {
	//panic("Not Implemented")
	total := 0.0
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func third() {
	fmt.Println("3rd")
}
func fourth() {
	fmt.Println("4th")
}

func fifth() {
	fmt.Println("5th")
}

func sum(xs ...int) int {
	total := 0
	for i := 0; i < len(xs); i++ {
		total += xs[i]
	}
	return total
}
