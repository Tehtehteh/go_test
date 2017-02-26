package main
import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main(){
	//defer fmt.Print("Defer func")
	//go fmt.Println("Hooyeta!")
	//for i:=0;i<10;i++{
	//	fmt.Printf("Hooyeta from loop %s\n", i)
	//}
	//defer f("deffered")
	go f("goroutine")
	f("direct")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}