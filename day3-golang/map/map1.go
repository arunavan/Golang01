package main  
import "fmt"  
func main ()  {  
   x := map[string]int{"Kate":28,"John":37, "Raj":20}  
   fmt.Print(x)  
   fmt.Println("\n",x["Raj"])  
}  

/*
package main  
import "fmt"  
func main() {  
   m := make(map[string]int)  
   fmt.Println(m)  
   m["Key1"] = 10  
   m["Key2"] = 20  
   m["Key3"] = 30  
   fmt.Println(m)  
   m["Key2"] = 555  
   fmt.Println(m)  
}  



package main  
import "fmt"  
func main() {  
   m := make(map[string]int)  
   m["Key1"] = 10  
   m["Key2"] = 20  
   m["Key3"] = 30  
   fmt.Println(m)  
   delete(m, "Key3")  
   fmt.Println(m)  
}     

package main  
import "fmt"  
func main() {  
   m := make(map[string]int)  
   m["Key1"] = 10  
   m["Key2"] = 20  
   m["Key3"] = 30  
   fmt.Println(m)  
   fmt.Println("The value:", m["Key2"])  
}  


package main  
import "fmt"  
func main() {  
   m := make(map[string]int)  
   m["Key1"] = 10  
   m["Key2"] = 20  
   m["Key3"] = 30  
   fmt.Println(m)  
   v, ok := m["Key2"]  
   fmt.Println("The value:", v, "Present?", ok)  
   i, j := m["Key4"]  
   fmt.Println("The value:", i, "Present?", j)  
}  


package main  
import "fmt"  
type Vertex struct {  
   Latitude, Longitude float64  
}  
var m = map[string]Vertex{  
   "JavaTpoint": Vertex{     40.68433, -74.39967,   },  
   "SSS-IT": Vertex{     37.42202, -122.08408,  },  
}  
func main() {  
   fmt.Println(m)  
}  
