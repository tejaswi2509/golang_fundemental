package main
  
import "fmt"
  
type Rectangle struct {
    Width  float64
    Height float64
}
type Circle struct{
	radius float64
}
type Shape interface {
    Area() float64
	
}
func (r Rectangle) Area() float64  {
    return r.Width * r.Height
}

func (c Circle) Area() float64  {
    return 3.14*c.radius * c.radius
}

func main(){
	var s Shape
	s = Rectangle{1,2}
	var c Shape
	c = Circle{2}
	fmt.Println("area of rectangle",s.Area())
	fmt.Println("area of circle",c.Area())

}