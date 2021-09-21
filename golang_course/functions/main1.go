package main
import "fmt"

func main(){
	//function called
	foo()
	bar("james")
	s1:= woo("moneypenny")
	fmt.Println(s1)
	x,y := mouse("Ian","Fleming")
	fmt.Println(x)
	fmt.Println(y)
	v:= sum(2,3,4,5,6,7,8,9)
	fmt.Println("the total is",v)

}
//function created
func foo(){
	fmt.Println("hello from foo")
}
func bar(s string){
	fmt.Println("hello",s)
}
func woo(st string) string{
	return fmt.Sprint("Hello from woo, ", st)
}
func mouse(fn string,ln string)(string, bool){
	a:= fmt.Sprint(fn," ",ln,`,says "hello"`)
	b:=true
	return a,b
}
//variadic parameter
func sum(x ...int)int{
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	sum:=0
	for i,v:=range x{
		sum+=v
		fmt.Println("for item in index position",i,"we are now adding",v,"to the total which is now",sum)
	}
	fmt.Println("the total sum is ",sum)
	return sum
}