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
	if(c<0){
		i--
	}
	return i
}
func powalg(a int,b int) int {
	var a2 int
	a2=a
	for c:=1;c<b;c++{
		a=a*a2
	}
	return a
}
func factoalg(a int) int{
	value:=1
	for c:=1;c<=a;c++{
		value=value*c
	}
	fmt.Println(value)
	return value
}
func primealg(a int) bool{
	var prime bool
	prime=true
	for i:=2;i<=a/2 && prime==true;i++{
		if (a%i==0){
			prime=false
		}
	}
	return prime
}

func main(){
	var a,b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println(a,"-",b,"=",subalg(a,b))
	fmt.Println(a,"*",b,"=",multalg(a,b))
	fmt.Println(a,"/",b,"=",divalg(a,b))
	fmt.Println(a,"^",b,"=",powalg(a,b))
	fmt.Println("!",a,"=",factoalg(a))
	fmt.Println("!",b,"=",factoalg(b))
	primoA:=primealg(a)
	primoB:=primealg(b)
	if (primoA==true) {
		fmt.Println(a,"es primo")
	}else{
		fmt.Println(a,"no es primo")
	}
	if (primoB==true) {
		fmt.Println(b,"es primo")
	}else{
		fmt.Println(b,"no es primo")
	}
}