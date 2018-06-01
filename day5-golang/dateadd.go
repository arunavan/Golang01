package main
  
import (
    "fmt"
    "time"
)
  
func main() {
    now := time.Now()
    fmt.Println("\nToday:", now)
      
    after := now.AddDate(1, 0, 0)
    fmt.Println("\nAdd 1 Year:", after)
      
    after = now.AddDate(0, 1, 0)
    fmt.Println("\nAdd 1 Month:", after)
      
    after = now.AddDate(0, 0, 1)
    fmt.Println("\nAdd 1 Day:", after)
      
    after = now.AddDate(2, 2, 5)
    fmt.Println("\nAdd multiple values:", after)
      
    after = now.Add(10*time.Minute)
    fmt.Println("\nAdd 10 Minutes:", after)
      
    after = now.Add(10*time.Second)
    fmt.Println("\nAdd 10 Second:", after)
      
    after = now.Add(10*time.Hour)
    fmt.Println("\nAdd 10 Hour:", after)
      
    after = now.Add(10*time.Millisecond)
    fmt.Println("\nAdd 10 Millisecond:", after)
      
    after = now.Add(10*time.Microsecond)
    fmt.Println("\nAdd 10 Microsecond:", after)
      
    after = now.Add(10*time.Nanosecond)
    fmt.Println("\nAdd 10 Nanosecond:", after)
}