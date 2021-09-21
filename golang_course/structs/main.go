package main
import "fmt"

type person struct{
	first string
	last string
	age int
}

type secretAgent struct{
	person
	ltk bool
}
func main(){
	sa1 := secretAgent{
		person: person{
			first: "james",
			last: "bond",
			age:32,
		},
		ltk:true,
	}
	p2 :=person{
		first:"miss",
		last:"moneypenny",
		age:27,
	}
	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(sa1.first,sa1.last,sa1.age,sa1.ltk)
	fmt.Println(p2.first,p2.last,p2.age)

	//anonymous struct
	s:= struct{
		first string
		friends map[string]int
		favDrinks []string
	}{
		first:"james",
		friends: map[string]int{
			"moneypenny":555,
			"Q":666,
			"M":777,
		},
		favDrinks: []string{
			"martini",
			"water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k,v:= range s.friends{
		fmt.Println(k,v)
	}
	for i,val:= range s.favDrinks{
		fmt.Println(i,val)
	}

}