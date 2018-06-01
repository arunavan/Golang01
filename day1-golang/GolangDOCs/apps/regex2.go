package main  
  
import (  
    "fmt"  
    "regexp"  
)  
  
func main() {  
    re := regexp.MustCompile(".com")  
    fmt.Println(re.FindStringIndex("google.com"))  
    fmt.Println(re.FindStringIndex("abc.org"))  
    fmt.Println(re.FindStringIndex("fb.com"))  
}  
