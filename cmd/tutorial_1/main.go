package main

import (
	"errors"
	"fmt"
	"strings"
)

type engine interface {
	interfacefunc() int
}

func (obj myStruct) structFunc() {
	fmt.Printf("name %v, age %v\n", obj.name, obj.age)
}

type myStruct struct {
	name string
	age  uint
}

// didnt understand interfaces
// func (obj engine){
//  fmt.Printf("name %v, age %v\n", obj.name, obj.age)
// }

// goroutine
// waitgroup
// channels
// generics

func genericFuc[T float32 | int | int32](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println("hello world")
	var intNum int // default value 0
	fmt.Println(intNum)
	var floatNum float32 = 1234.9
	fmt.Println(floatNum)
	// cant perform operation on diff types
	var str string = "dhruv"
	str = `raw
    string`
	str = "dhruv"
	fmt.Println(str)
	fmt.Println(len(str))
	var inferedtype = "dhruv"
	fmt.Println(inferedtype)
	xyz := "dhruv"
	fmt.Println(xyz)
	const myconst string = "cant change"
	// ans, err := myfunc(10, 2)
	// if err != nil {
	//  print(err.Error())
	// } else {
	//  print(ans)
	// }

	// var intarr [2]int = [2]int{1, 2}
	intarr := [...]int{1, 2}
	println(intarr[0])
	// slice
	fmt.Println(intarr[0:1])
	// omite length value
	slice := []int{1, 2, 3}
	slice = append(slice, 10)
	fmt.Println(slice)
	// capacity
	println(cap(slice))
	slice = append(slice, 10)
	println(cap(slice))
	// make(type, size, capacity)
	var newslice []int = make([]int, 3, 10)
	fmt.Println(newslice)
	var mymap = map[int]int{1: 2}
	delete(mymap, 1)
	var v, exist = mymap[1]
	if exist {
		fmt.Println(v)
	} else {
		fmt.Println("doesnt exist")
	}
	fmt.Println(mymap[1999]) // return default value
	mymap[10] = 20
	mymap[30] = 40
	// order is not preserved
	for name := range mymap {
		fmt.Printf("key :%v val:%v\n", name, mymap[name])
	}
	for idx, val := range intarr {
		fmt.Printf("idx :%v val:%v\n", idx, val)
	}
	// string ar immutable in go
	var strSlice = []string{"d", "h", "r", "u", "v"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Println(catStr)
	var mystructobj myStruct = myStruct{name: "dhruv", age: 10}
	fmt.Println(mystructobj)
	mystructobj.structFunc()
	fmt.Println(genericFuc[int](1, 2))
	fmt.Println(genericFuc[float32](1.1, 2.0))
}

// general pattern to return error with the answer
func myfunc(a int, b int) (int, error) {
	var err error
	if b == 0 {
		err = errors.New("divide by zero")
		return 0, err
	}
	return a / b, err
}
