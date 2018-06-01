package main  
  
import (  
    "fmt"  
    "regexp"  
)  
  
func main() {  
    re := regexp.MustCompile(".com")  
    fmt.Println(re.FindString("google.com"))  
    fmt.Println(re.FindString("abc.org"))  
    fmt.Println(re.FindString("fb.com"))  
}  
