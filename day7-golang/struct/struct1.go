package main  
import (  
   "fmt"  
)  
type person struct {  
   firstName string  
   lastName  string  
   age       int  
}  
func main() {  
   x := person{age: 30, firstName: "John", lastName: "Anderson", }  
   fmt.Println(x)  
   fmt.Println(x.firstName)  
}  
