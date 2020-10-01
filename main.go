package main

import "fmt"

func main() {
	var b, h int
	fmt.Scanln(&b)
	fmt.Scan(&h)
	area := (b * h) / 2
	fmt.Println(area)
}
