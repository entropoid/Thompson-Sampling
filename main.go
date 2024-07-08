package main

import (
  "fmt"
)




func main()  {
  // Say we have some restaurants listed on our platform and we want to recommend users most apt restaurant
  // for their next dining experience. 

  // we want to come up with a mechanism to recommend options to users such that new options are not 
  // explicitly kept out. i.e. there is a trade-off between exploration v/s exploitation.

  // In this implmentation, I want to assume that we have 4 restaurants, and initially all are beta distri.
  // in exploration stage i.e. no historical data is available.
  var num_restaurants int = 4
  var num_trials int = 20

  var ts *ThompsonSampling = InitThompsonSampling(num_restaurants)
  
  // run trials
  for trial:=0;trial<num_trials;trial++ {
    fmt.Println("Trial ",trial)

    // get Beta samples
    var samples []float64 = ts.Sample()
    
    fmt.Println(samples)
  
    // get suggested_restaurant
    var suggested_restaurant int = ts.Choose(samples)
    fmt.Println("Suggested restaurant: ", suggested_restaurant)
  
    // feedback
    fmt.Println("How was the experience? (good/bad)")
    var feedback string
    fmt.Scan(&feedback)
    if feedback == "good" {
      ts.Feedback(suggested_restaurant,true)
    } else {
      ts.Feedback(suggested_restaurant,false)
    }

    fmt.Println("Updated values of params: ")
    fmt.Println(ts.restaurants)
  }

}
