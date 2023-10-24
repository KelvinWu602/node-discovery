package main

import ("fmt"
		"github.com/KelvinWu602/node-discovery/sidanLa"
		)

func main() {
	fmt.Println("Hello World")
	result := sidanLa.Function1(1,1) //type inference
	fmt.Println(result)
}