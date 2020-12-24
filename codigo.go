package main

import (
	"fmt"
)

func subalg(a int , b int) int {
	var c int
	
	if b>a {
		for c=0;a<b;c++ {
			a++
		}
	}else{
		for c=0;b<a;c++ {
			b++
		}
	}
	return c
}
func multalg(a int,b int) int {
	c:= a
	for i:=1;i<b;i++{
		c=c+a
	}
	return c
}
func divalg(a int,b int) int{
	c:=a
	var i int
	for i=0;c>0;i++{
		c=c-b
	}
	fmt.Println(i)
	return i
}

func main(){
	var a,b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println(a,"-",b,"=",subalg(a,b))
	fmt.Println(a,"*",b,"=",multalg(a,b))
	divalg(a,b)
}