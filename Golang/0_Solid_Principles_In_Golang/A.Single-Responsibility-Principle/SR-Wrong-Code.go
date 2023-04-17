package main
import(
	"fmt"
	"math"
)

type Circle struct{
	radius float32
}

func(c Circle) Area(){
	fmt.Println("Circle Area: ",math.Pi*c.radius*c.radius) //Not Following Single Responsility Principle.
}

type Square struct{
	sideLen float32
}

func(s Square) Area(){
	fmt.Println("Circle Area:",s.sideLen*s.sideLen)//Not Following Single Responsility Principle.
}

//Driver Code
func main(){
	//Circle Object.
	areaObject := Circle{radius:5}
	areaObject.Area()

	//Square Object.
	circleObject := Square{sideLen: 56 }
	circleObject.Area()
}
/*
$ go run SR-Wrong-Code.go
Circle Area:  78.53982
Circle Area: 3136
*/