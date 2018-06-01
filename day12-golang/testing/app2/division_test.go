package gotest

import (
    "testing"
)

func Test_Division_1(t *testing.T) {
    // try a unit test on function
    if i, e := Division(6, 2); i != 3 || e != nil { 
        // If it is not as expected, then the test has failed 
        t.Error("division function tests do not pass ") 
    } else {
        // record the expected information
        t.Log("....first test passed ",e) 
    }
}
//comment below code  and 
//run with go test -v
//func Test_Division_2(t *testing.T) {
 //   t.Error("2nd method just does not pass")
  
//}