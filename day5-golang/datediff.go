package main
  
import (
         "fmt"
         "time"
)
  
func main() {
loc, _ := time.LoadLocation("UTC")
now := time.Now().In(loc)
fmt.Println("\nToday : ", loc, " Time : ", now)
  
pastDate := time.Date(2015, time.May, 21, 23, 10, 52, 211, time.UTC)
fmt.Println("\n\nPast  : ", loc, " Time : ", pastDate) // 
fmt.Printf("###############################################################\n")
diff := now.Sub(pastDate)
  
hrs := int(diff.Hours())
fmt.Printf("Diffrence in Hours : %d Hours\n", hrs)
  
mins := int(diff.Minutes())
fmt.Printf("Diffrence in Minutes : %d Minutes\n", mins)
  
second := int(diff.Seconds())
fmt.Printf("Diffrence in Seconds : %d Seconds\n", second)
  
days := int(diff.Hours() / 24)
fmt.Printf("Diffrence in days : %d days\n", days)
  
fmt.Printf("###############################################################\n\n\n")
  
futureDate := time.Date(2018, time.June, 21, 23, 10, 52, 211, time.UTC)
fmt.Println("Future  : ", loc, " Time : ", futureDate) // 
fmt.Printf("###############################################################\n")
diff = futureDate.Sub(now)
  
hrs = int(diff.Hours())
fmt.Printf("Diffrence in Hours : %d Hours\n", hrs)
  
mins = int(diff.Minutes())
fmt.Printf("Diffrence in Minutes : %d Minutes\n", mins)
  
second = int(diff.Seconds())
fmt.Printf("Diffrence in Seconds : %d Seconds\n", second)
  
days = int(diff.Hours() / 24)
fmt.Printf("Diffrence in days : %d days\n", days)
  
}