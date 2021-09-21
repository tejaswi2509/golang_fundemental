package main
import "fmt"

func main(){
	s := []int{4,5,7}

	for i,v:= range s{
		fmt.Println(i,v)
	}

	//slicing
	fmt.Println(s[1:])
	fmt.Println(s[2:3])
	fmt.Println(s[:3])

	t:=[]int{4,5,6}
	//append slice
	t = append(t,77,88,99,1014)
	fmt.Println(t)
	q:=[]int{234,456,678,987}
	t=append(t,q...)
	fmt.Println(t)

	// deleting 
	t=append(t[:2],t[4:])
	fmt.Println(t)

}