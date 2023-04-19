/*
The Dependency Inversion Principle states, high-level modules or classes should not depend on low-level modules or classes, but both should depend on abstractions. In other words, 
the principle suggests that the dependency direction should be inverted.

In practical terms, this means that instead of having a high-level module or class directly use a low-level module or class, both should depend on an abstraction or an interface. 
This allows for greater flexibility in the code, as changes to the low-level module or class do not impact the high-level module or class, as long as the interface remains the 
same.

---Pros of DIP---
A. The DIP promotes loose coupling between modules and classes, which can make the code more maintainable, testable, and extensible. 
B. It also promotes the use of interfaces, which allows for different implementations of the same interface to be swapped out without affecting the rest of the code.
*/

package main
import(
	"fmt"
	"log"
)

type DBConnection interface{
	QueryData() []string
}

//DB Connection
type MySQL struct{

}

func(db MySQL) QueryData() ([]string){
	return []string{"Object-1","Object-2","Object-3"}
}

//Instead of having dbObject of type MySQL we've declared it of type interface i.e. DBConnection.
type Repository struct{
	dbObject DBConnection
}

func(r Repository) GetData() []string{
	return r.dbObject.QueryData()
}

// Begins: Dummy models to verify the correctness of functionality.
type Dummy struct{}

func(d Dummy) DummyMethod(){
	log.Println("DummyMethod Awake!!!")
}
//Ends: Dummy models to verify the correctness of functionality.
func main(){
	mySQLObject := new(MySQL)
	repoObject := Repository{dbObject: *mySQLObject}
	fmt.Println(repoObject.GetData())

	//Below code will not work as Dummy struct as not implemented DBConnection interface like QueryData
	/*
	dummyObject := new(Dummy)
	repoObject = Repository{dbObject: *dummyObject}
	fmt.Println(repoObject.GetData())
	*/
	//---Below is the output when we uncomment above code.
	/*
	$ go run DIP-Correct-Code.go 
	# command-line-arguments
	.\DIP-Correct-Code.go:55:36: cannot use *dummyObject (variable of type Dummy) as DBConnection value in struct literal: Dummy does not implement DBConnection (missing method QueryData)
	*/
}

/*
Output:
$ go run DIP-Correct-Code.go 
[Object-1 Object-2 Object-3]
*/