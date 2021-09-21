package main
import "fmt"

// declaring the value in package level
var k int

// creating own type of variable
type hotdog int
var b hotdog
var a int

func main(){
	//using var keyword
	var i int
	i=1
	//using short declaration
	j:="s"
	// giving value to package level variable
	k = 3
	// zero value
	var z int
	// defining a value for variable of type hotdog
	b =43
	a=4
	fmt.Println(a)
	//type conversion
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	fmt.Println(b)
	fmt.Printf("%T\n",b)
	fmt.Println(i,j,k,z)
}