package main
import (
    "log"
    "os"
)
func main() {
    path:="test.txt"
    newpath:="test2.txt"
    err :=os.Rename(path,newpath)
    if err != nil {
        log.Fatal(err)
    }
}

