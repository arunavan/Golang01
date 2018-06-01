package main  
import (  
   "os"  
   "log"  
   "io/ioutil"  
   "fmt"  
)  
func main() {  
   file, err := os.Create("file.txt")  
   if err != nil {  
      log.Fatal(err)  
   }  
   file.WriteString("Hi... there")  
   file.Close()  
   stream, err:= ioutil.ReadFile("file.txt")  
   if err != nil {  
      log.Fatal(err)  
   }  
   readString := string(stream)  
   fmt.Println(readString)  
}  
