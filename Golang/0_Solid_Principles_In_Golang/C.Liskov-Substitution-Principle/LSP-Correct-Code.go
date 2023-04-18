/*
Liskov Substitution Principle states that objects of a superclass should be replaceable with objects of it's subclasses without breaking the application.
In other words, objects of subclasses should behave same way as the object of superclass.
*/

package main
import(
	"fmt"
	"reflect"
	"log"
)

type Transport interface{
	GetName() string
}

//A sort of common struct which will have Vehicle name.
type Vehicle struct{
	name string
}

func(v Vehicle) GetName() string{
	return v.name
}

//As we're defining car type of vehicle we're using Composite inheritence where we're embedding "Vehicle" into Car.
type Car struct{
	Vehicle 
	wheel int
	doors int
}

type Motorcycle struct{
	Vehicle
	wheel int
}

type Printer struct{}


func(printer Printer) printTransportName(p Transport){
	log.Println("Type of object: ",reflect.TypeOf(p))
	fmt.Println("Name: ",p.GetName())
}
/*
In below we're passing different object but output is not getting impacted. 
Hence code is following "Liskov Substitution Principle". 
*/
func main(){
	hummerVehicleObject := new(Vehicle)
	hummerVehicleObject.name = "Hummer"
	
	carHummerObject := Car{Vehicle: *hummerVehicleObject, wheel: 4,doors: 4}
	motorcycleObject := Motorcycle{
		Vehicle: Vehicle{name: "Kawasaki"},
		wheel : 2,
	}
	printerObject := new(Printer)
	//Passing Motorcycle object.
	
	//Passing Car struct object.
	printerObject.printTransportName(carHummerObject)
	//Passing Motorbyke object.
	printerObject.printTransportName(motorcycleObject)
	//Passing Hummer struct object.
	printerObject.printTransportName(hummerVehicleObject)

}
/*
$ go run LSP-Correct-Code.go 
2023/04/18 16:30:29 Type of object:  main.Car
Name:  Hummer
2023/04/18 16:30:29 Type of object:  main.Motorcycle
Name:  Kawasaki
2023/04/18 16:30:29 Type of object:  *main.Vehicle
Name:  Hummer
*/