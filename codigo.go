package main

import (
	"fmt"
)

func main(){
	var a int
	var b int 
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	var c int
	
	if b>a {
		for c=0;a<b;c++ {
			a++
		}
		fmt.Println("la diferencia es de ",c)
	}else{
		for c=0;b<a;c++ {
			b++
		}
		fmt.Println("la diferencia es de ",c)
	}
}