package main
import "fmt"

func main(){
	a:=43
	fmt.Println(a)
	fmt.Println(&a)

	var b = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 42
	fmt.Println(b)

	//without pointer
	t:=0
	fo(t)
	fmt.Println(t)
	//with pointer
	x:=0
	fmt.Println("x before",&x)
	fmt.Println("x before",x)
	foo(&x)
	fmt.Println("x after",&x)
	fmt.Println("x after",x)

}

func fo(y int){
	fmt.Println(y)
	y=43
	fmt.Println(y)
}

func foo(y *int){
	fmt.Println("y before",y)
	fmt.Println("y before",*y)
	*y =43
	fmt.Println("y after",y)
	fmt.Println("y after",*y)
}