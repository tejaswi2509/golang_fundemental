package main
import "fmt"

func main(){
	for i:=33; i<=36;i++{
		if i%2!=0{
			fmt.Printf("%v\t#%X\t#%U\n",i,i,i)
		}else if i==41{
			fmt.Println("our value was 41")
		}else{
			fmt.Println("our value was not 41")
		}
	}

	n:="bond"
	switch n{
	case "moneypenny","bond":
		fmt.Println("moneypenny or bond")
	case "M":
		fmt.Println("M")
	case "Q":
		fmt.Println("Q")
	default:
		fmt.Println("this is default")
	}
	fmt.Printf("true && false\t %v\n", true&&false)

}