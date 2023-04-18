package main
import(
	"fmt"
	"math"
	"log"
)

var(
	sum float64
)

type Shape interface{
	Area() float64
}

type Circle struct{
	radius float64
}

func(c Circle) Area() float64{
	log.Println("Circle: ",math.Pi * c.radius * c.radius)
	return math.Pi * c.radius * c.radius
} 

type Square struct{
	sideLen float64
}

func(s Square) Area() float64{
	log.Println("Square: ",s.sideLen * s.sideLen)
	return s.sideLen * s.sideLen
}

type Triangle struct{
	height float64
	base float64
}

func(t Triangle) Area() float64{
	log.Println("Triangle: ",t.base * t.height / 2)
	return t.base * t.height / 2
}

type Calculator struct{
	sum float64
}
/*
Instead of using switch case we're calling respective methods as there objects are passed.
Say if we've a new requirement like Triangle then we can simply extend it no need to tweak the existing functionality.
*/
func(c *Calculator) FindTotalArea(shapes ... Shape) float64{
	for _,shape := range shapes{
		sum += shape.Area()
	}
	return sum
}

func main(){
	circleObject := new(Circle)
	circleObject.radius = 5

	squareObject := Square{sideLen: 5}

	triangleObject := new(Triangle)
	triangleObject.height = 2
	triangleObject.base = 2

	calcObject := new(Calculator)

	fmt.Println("Total Area: ",calcObject.FindTotalArea(*circleObject,squareObject,*triangleObject))

}

/*
Output:
$ go run OC-Correct-Code.go 
2023/04/18 11:30:44 Circle:  78.53981633974483
2023/04/18 11:30:44 Square:  25
2023/04/18 11:30:44 Triangle:  2
Total Area:  105.53981633974483
*/