package main
import (
    "log"
    "os"
)
func main() {
    file,err :=os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    file1,err :=os.OpenFile("test.txt",os.O_APPEND,0666)
    if err != nil {
        log.Fatal(err)
    }
    file1.Close()

    //os.O_RDONLY,os.O_WRONLY,os.O_RDWR,os.O_CREATE
    //os.O_TRUNC,os.O_APPEND

}

