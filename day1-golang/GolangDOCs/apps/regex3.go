package main  
  
import (  
    "fmt"  
    "regexp"  
)  
  
func main() {  
    re := regexp.MustCompile("f([a-z]+)ing")  
    fmt.Println(re.FindStringSubmatch("flying"))  
    fmt.Println(re.FindStringSubmatch("abcfloatingxyz"))  
}  
