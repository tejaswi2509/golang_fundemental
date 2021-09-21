package main
import "fmt"

func main(){
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
	y:=[2]int{2,1}
	fmt.Println(y)

}