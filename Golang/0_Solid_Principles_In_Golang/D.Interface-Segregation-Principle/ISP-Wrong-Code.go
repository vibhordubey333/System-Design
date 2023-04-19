/*
Interface Segregation Principle states, clients should not be forced to depend on methods they don't use.
In simple terms, "Keep interfaces simple, prefrably just one method"
*/
package main

import (
	"fmt"
	"math"
	"reflect"
	"log"
)

type Shape interface {
	Area() float64
	Volume() float64
}

type Square struct {
	sideLen float64
}

func (s Square) Area() float64 {
	return s.sideLen * s.sideLen
}

// This method should be present in separate interface as Square struct should not be forced to implement it as it is not needed.
// If we comment it we'll get the error.
func (s Square) Volume() float64 {
	return 0
}

type Cube struct {
	sideLen float64
}

func (c Cube) Area() float64 {
	return math.Pow(c.sideLen, 2)
}

func (c Cube) Volume() float64 {
	return math.Pow(c.sideLen, 3)
}

// Sum the shapes area
func AreaSum(shapes []Shape) float64 {
	
	var sum float64
	for _, s := range shapes {
		log.Println("Type Received:",reflect.TypeOf(s))
		sum += s.Area()
	}
	return sum
}

func main() {

	squareObject := Square{sideLen: 5}
	cubeObject := Cube{sideLen: 3}
	shapeObject := []Shape{squareObject,cubeObject}
	fmt.Println(AreaSum(shapeObject))
}
