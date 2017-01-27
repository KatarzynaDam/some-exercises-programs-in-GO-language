package main

import "fmt"
import "time"

func main()  {
  
  fmt.Println("Kasia_" + "Dam_")
  
  a := fmt.Println
  now := time.Now()  //pobieram aktualny czas // I'm getting curently time
  a(now)
  
  then := time.Date(2017, 1, 26, 20, 50, time.UTC )
  a(then)
  
  a(then.Year)
  a(then.Month)
  a(then.Day)
  a(then.Hour)
  a(then.Minute)
  a(then.Second)
  a(then.Weekday)
  
  a(then.After(now))
  a(then.Equal(now))
  a(then.Before(now))
  
  other := now.Sub(then)  // other -> it's a different time; I'm comparing the time now and thime "other"
  a(other)                // other -> zmienna z innym czasem (innym niż zadeklarowane powyżej "now"); porównuję te dwa czasy
  
  a(other.Seconds())
  a(other.Minutes())
  a(other.Hours())
  a(other.Days())
  
  a(then.Add(other))
  a(then.Add(-other))
  
  
}
