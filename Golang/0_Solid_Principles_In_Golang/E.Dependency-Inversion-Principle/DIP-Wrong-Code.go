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
)

//DB Connection
type MySQL struct{

}

func(db MySQL) QueryData() ([]string){
	return []string{"Object-1","Object-2","Object-3"}
}

type Repository struct{
	dbObject MySQL
}

func(r Repository) GetData() []string{
	return r.dbObject.QueryData()
}

func main(){
	mySQLObject := new(MySQL)
	repoObject := Repository{dbObject: *mySQLObject}
	fmt.Println(repoObject.GetData())
}