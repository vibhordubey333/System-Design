//Open-Close principle states, a software artifact should be open for extension but closed for modifications.

package main
import(
	"fmt"
	"math"
	"log"
	"reflect"
)

var(
	sum float64
)


type Circle struct{
	radius float64
}

type Square struct{
	sideLen float64
}

type Calculator struct{
	total float64
}

//FindArea method doesn't follow the Open-Close principle. Say if we've a new type like Triangle then we need to modify the code again which is not correct approach.
func(c Calculator) FindArea(shapes...interface{}) float64{
	
	for _,shape := range shapes{
		//log.Println("Type: ",reflect.TypeOf(shape))
		switch shape.(type){
		case Circle:
			r := shape.(Circle).radius
			log.Println("Circle: ",math.Pi * r * r)
			sum += math.Pi * r * r
			
		case Square:
			l := shape.(Square).sideLen
			log.Println("Square: ", l * l)
			sum += l * l

		default:
				fmt.Println("Invalid Type")
		}
	}

	return sum
}

func main(){
	//When using new() function you need to dereference the variable.
	circleObject := new(Circle)//Circle{radius: 5}
	//When manually creating object you don't need to dereference the variable.
	squareObject := Square{sideLen: 10}//new(Square)

	calculatorObject := Calculator{}
	circleObject.radius = 5
	squareObject.sideLen = 10
	 
	fmt.Println("Total Of Areas: ",calculatorObject.FindArea(*circleObject,squareObject))
}

/*
$ go run OC-Wrong-Code.go
2023/04/18 11:39:42 Circle:  78.53981633974483
2023/04/18 11:39:42 Square:  100
Total Of Areas:  178.53981633974485
*/