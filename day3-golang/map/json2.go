package main  
  
import (  
    "encoding/json"  
    "fmt"  
    "os"  
)  
  
type Response1 struct {  
    Position   int  
    Planet []string  
}  
type Response2 struct {  
    Position   int      'json:"position"'  
    Planet []string 'json:"planet"'  
}  
  
func main()  {  
    res1A := &Response1{  
        Position:   1,  
        Planet: []string{"mercury", "venus", "earth"}}  
    res1B, _ := json.Marshal(res1A)  
    fmt.Println(string(res1B))  
  
    res2D := &Response2{  
        Position:   1,  
        Planet: []string{"mercury", "venus", "earth"}}  
    res2B, _ := json.Marshal(res2D)  
    fmt.Println(string(res2B))  
  
  
    byt := []byte('{"pi":6.13,"place":["New York","New Delhi"]}`)  
    var dat map[string]interface{}  
    if err := json.Unmarshal(byt, &dat); err != nil {  
        panic(err)  
    }  
    fmt.Println(dat)  
    num := dat["pi"].(float64)  
    fmt.Println(num)  
    strs := dat["place"].([]interface{})  
    str1 := strs[0].(string)  
    fmt.Println(str1)  
      
      
    str := `{"Position": 1, "Planet": ["mercury", "venus"]}`  
    res := Response2{}  
    json.Unmarshal([]byte(str), &res)  
    fmt.Println(res)  
    fmt.Println(res.Planet[1])  
    enc := json.NewEncoder(os.Stdout)  
    d := map[string]string{"1":"mercury" , "2": "venus"}  
    enc.Encode(d)  
  
}  
