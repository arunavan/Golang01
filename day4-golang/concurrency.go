package main  
import (  
   "fmt"  
   "time"  
   "sync"  
)  
var wg = sync.WaitGroup{}  

  //routine1 is function which is invoked by go command
  // all routines are lightweight and sharing common addressspace

func main() {  
   wg.Add(2)  
   go fun1()   //
   go fun2()    //routine2
   wg.Wait()  
}  
func fun1(){  
   for  i:=0;i<10;i++{  
      fmt.Println("fun1,  ->",i)  
      time.Sleep(time.Duration(5*time.Millisecond))  
   }  
   wg.Done()  
}  
func fun2(){  
   for i:=0;i<10;i++{  
      fmt.Println("fun2,  ->",i)  
      time.Sleep(time.Duration(10*time.Millisecond))  
   }  
   wg.Done()  
}  
