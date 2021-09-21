package main

import "fmt"

func main(){
	m :=map[string]int{
		"james":32,
		"Miss Moneypenny":27,
	}
	fmt.Println(m)
	fmt.Println(m["james"]) //getting data using key
	fmt.Println(m["barnabas"])

	v,ok:=m["barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	// adding data to map
	m["todd"] = 33

	if v,ok:= m["barnabas"]; ok{
		fmt.Println(v)
	}

	for k,v:=range m{
		fmt.Println(k,v)
	}
	//deleting
	delete(m,"james")
	fmt.Println(m)

}