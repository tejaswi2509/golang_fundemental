package main
import "fmt"
func main(){
	//anonymous function
	func(){
		fmt.Println("anonymous function ran")
	}()

	func(x int){
		fmt.Println("the value of x is ",x)
	}(42)
	//function can be assigned to a variable
	f := func(){
		fmt.Println("func to f variable")
	}
	f()

	g := func(x int){
		fmt.Println("the year brother started watching:",x)
	}
	g(1998)

	//function to return a function
	fmt.Println(bar()())

}

func foo() string{
	s:="hello"
	return s
}
func bar()func()int{
	return func()int{
		return 451
	}
}