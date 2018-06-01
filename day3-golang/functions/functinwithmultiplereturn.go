package main  
import (  
   "fmt"  
)  
func main() {  
   fmt.Println(addAll(10,15,20,25,30))  
}  
func addAll(args ... int)(int,int)  {  
   finalAddValue:=0  
   finalSubValue:=0  
   for _,value  := range args{  
      finalAddValue += value  
      finalSubValue -= value  
   }  
   return finalAddValue,finalSubValue  
} 
