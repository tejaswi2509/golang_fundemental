package main
import (
 "fmt"
 "runtime"
 
)
var f bool

func main() {
	fmt.Println("goroutines : ",runtime.NumGoroutine())
 	for {
		go anotherGoroutine()
	 }
	
	forever()
	
	fmt.Println("goroutines : ",runtime.NumGoroutine())
 	
}
func forever() {
  for {
	
  }

}
func anotherGoroutine() {
 i := 0
 for {
  i++
  fmt.Print(i)
  fmt.Print("\n")
 }
}