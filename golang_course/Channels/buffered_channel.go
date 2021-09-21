package main 
import "fmt"

func main(){
	c := make(chan int,2)
	c <- 42
	c <- 24
	fmt.Println(<-c)
	fmt.Println(<-c)
}