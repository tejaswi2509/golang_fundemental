package main
import ("fmt"
	// "time"
)

func main(){
	defer t()
	
	fmt.Println("main ended")
	c()
	
	// time.Sleep(time.Second * 10)

}

func c(){
	fmt.Println("c function")
}
func t(){
	go g()
	
}
func g(){
	fmt.Println("goroutine")
}