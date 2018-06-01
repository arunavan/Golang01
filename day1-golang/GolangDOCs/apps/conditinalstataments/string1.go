package main  
import ("fmt"  
      "reflect"  
)  
func main()  {  
   var x string = "Hello World"  
   fmt.Println(x)  
   fmt.Println(reflect.TypeOf(x))  
}  


/*

package main  
import "fmt"  
func main() {  
   fmt.Println("Ascii value of A is ","A"[0])  
}  



package main  
import "fmt"  
import "strings"  
func main() {  
   str := "india"  
   fmt.Println(strings.ToUpper(str))  
}  


package main  
import "fmt"  
import "strings"  
func main() {  
   str := "INDIA"  
   fmt.Println(strings.ToLower(str))  
} 

package main  
import "fmt"  
import "strings"  
func main() {  
   s := "INDIA"  
   fmt.Println(strings.HasPrefix(s,"IN"))  
}  

package main  
import "fmt"  
import "strings"  
func main() {  
   s := "INDIA"  
   fmt.Println(strings.HasSuffix(s,"IA"))  
}  


package main  
import "fmt"  
import "strings"  
func main() {  
   var arr = []string{"a","b","c","d"}  
   fmt.Println(strings.Join(arr,"*"))  
}  

package main  
import "fmt"  
import "strings"  
func main() {  
   str:= "Hi...there"  
   fmt.Println(strings.Contains(str,"th"))  
}  

package main  
import "fmt"  
import "strings"  
func main() {  
   str:= "Hi...there"  
   fmt.Println(strings.Index(str,"th"))  
} 


package main  
import "fmt"  
import "strings"  
func main() {  
   str:= "Hi...there"  
   fmt.Println(strings.Count(str,"e"))  
}  

package main  
import "fmt"  
import "strings"  
func main() {  
   str:= "Hi...there"  
   fmt.Println(strings.Replace(str,"e","Z",2))  
}  

package main  
import "fmt"  
import "strings"  
func main() {  
   str := "I,love,my,country"  
   var arr []string = strings.Split(str, ",")  
   fmt.Println(len(arr))  
   for i, v := range arr {  
      fmt.Println("Index : ", i, "value : ", v)  
   }  
} 

 
package main  
import (  
   "fmt"  
   "strings"  
)  
func main() {  
   fmt.Printf("%q\n", strings.Split("x,y,z", ","))  
   fmt.Printf("%q\n", strings.Split(" John and Jack and Johnny and Jinn ", "and"))  
   fmt.Printf("%q\n", strings.Split(" abc ", ""))  
   fmt.Printf("%q\n", strings.Split("", "Hello"))  
}  


package main  
import (  
   "fmt"  
   "strings"  
)  
func main() {  
   fmt.Println(strings.Compare("a", "b"))  
   fmt.Println(strings.Compare("a", "a"))  
   fmt.Println(strings.Compare("b", "a"))  
}  


package main  
import (  
   "fmt"  
   "strings"  
)  
func main() {  
      fmt.Println(strings.TrimSpace(" \t\n I love my country  \n\t\r\n"))  
}  



package main  
import (  
   "fmt"  
   "strings"  
)  
func main() {  
   fmt.Println(strings.ContainsAny("Hello", "A"))  
   fmt.Println(strings.ContainsAny("Hello", "o & e"))  
   fmt.Println(strings.ContainsAny("Hello", ""))  
   fmt.Println(strings.ContainsAny("", ""))  
}  







*/