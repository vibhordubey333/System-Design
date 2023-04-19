/*
Interface Segregation Principle states, clients should not be forced to depend on methods they don't use.
In simple terms, "Keep interfaces simple, prefrably just one method"

*/
package main

import (
	"fmt"
	"reflect"
	"log"
	"math"
)

type Shape interface {
	Area() float64
}
/*
//Yet to ascertain the usecase of it.
type Object interface {
	Shape
	Volume() float64
}
*/

type Square struct {
	sideLen float64
}

func (s Square) Area() float64 {
	return s.sideLen * s.sideLen
}

type Cube struct{
	// --- old code 
	//sideLen float64
	Square // Instead of repeating the fields here we can use Square itself as sideLen is only we need & it is present in Square.
}

func(c Cube) Area() float64{
	return math.Pow(c.sideLen,2)
}

func(c Cube) Volume() float64{
	return math.Pow(c.sideLen,3)
}
//// As a part of correct code commenting this method shouldn't give any error.
// This method should be present in separate interface as Square struct should not be forced to implement it as it is not needed.
// If we comment it we'll get the error.
/*func (s Square) Volume() float64 {
	return 0
}*/

// Sum the shapes area
func AreaSum(shapes []Shape) float64 {
	var sum float64
	for _, s := range shapes {
		log.Println("Type Recieved: ",reflect.TypeOf(s))
		sum += s.Area()
	}
	return sum
}

func main() {

	squareObject := Square{sideLen: 5}
	cubeObject := Cube{Square: Square{sideLen:3}}
	shapeObject := []Shape{squareObject,cubeObject}

	fmt.Println(AreaSum(shapeObject))
}
/*
$ go run ISP-Correct-Code.go 
2023/04/19 11:39:20 Type Recieved:  main.Square
2023/04/19 11:39:20 Type Recieved:  main.Cube
34
*/
