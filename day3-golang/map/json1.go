package main  
import "encoding/json"  
import "fmt"  
  
func main() {  
    bolType, _ := json.Marshal(false) //boolean Value  
    fmt.Println(string(bolType))  
    intType, _ := json.Marshal(10) // integer value  
    fmt.Println(string(intType))  
    fltType, _ := json.Marshal(3.14) //float value  
    fmt.Println(string(fltType))  
    strType, _ := json.Marshal("JavaTpoint") // string value  
    fmt.Println(string(strType))  
    slcA := []string{"sun", "moon", "star"} //slice value  
    slcB, _ := json.Marshal(slcA)  
    fmt.Println(string(slcB))  
    mapA := map[string]int{"sun": 1, "moon": 2} //map value  
    mapB, _ := json.Marshal(mapA)  
    fmt.Println(string(mapB))  
}  
