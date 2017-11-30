package main

import("fmt"
       "exercises/fizzbuzz/fb")


func main(){

	n := 1

	for n <= 100 {
		fmt.Println(fb.Fizzbuzz(n))
		n++
	}
}
