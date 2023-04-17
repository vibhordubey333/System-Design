package main
import(
	"fmt"
	"math"
)
/*
Reference:
1. https://faun.dev/c/stories/josueparra2892/golang-solid-principles/
*/

/*
Single Responsibility Principle:

1. Modify methods to perform single functionality. Don't print as we're doing in wrong code.
2. Define a new type of handle to print the area output. We will defining a new interface "Shape". Then add "Area" method to interface "Shape". Then declare struct variable and implement "ToText()" method with it.
3. Instead of printing the variable pass the "Circle" and "Square" object to 

*/

type Shape interface{
	Area() float32
}


type OutPrinter struct{}

func(op OutPrinter) ToText(s Shape) string{
	return fmt.Sprintf("The area is: %f",s.Area())
}

type Circle struct{
	radius float32
}

func(c Circle) Area() float32{
	return math.Pi * c.radius * c.radius // Single Responsility Principle.
}

type Square struct{
	sideLen float32
}

func(s Square) Area() float32{
	return s.sideLen*s.sideLen// Single Responsility Principle.
}

//Driver Code
func main(){
	objectOutPrinter := OutPrinter{}
	//Circle Object.
	areaObject := Circle{radius:5}
	areaObject.Area()
	fmt.Println("Circle Area: ",objectOutPrinter.ToText(areaObject))

	//Square Object.
	circleObject := Square{sideLen: 56 }
	circleObject.Area()
	fmt.Println("Square Area: ",objectOutPrinter.ToText(circleObject))
}

/*
$ go run SR-Correct-Code.go
Circle Area:  78.53982
Circle Area: 3136
*/