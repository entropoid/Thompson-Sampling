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
  // with historical context with weighted decay
  var num_restaurants int = 4
  var num_trials int = 20
  var decay_factor float64 = 1 // has to be between 0 and 1; 1 means non-contextual TS 
  // different modelling of decay factor can give different results.
  var history_choices []int = []int{1,1,2,3,2,2,0,0,2,2}
  var feedback []int = []int{1,1,0,0,1,1,0,0,0,0}
  


  var ts *ThompsonSampling = ContextualThompsonSampling(history_choices,feedback,num_restaurants,decay_factor)
  
  // run trials
  for trial:=0;trial<num_trials;trial++ {
    fmt.Println("Trial ",trial)
    fmt.Println("values of params: ")
    fmt.Println(ts.restaurants)


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

  }
  fmt.Println("Updated values of params: ")
  fmt.Println(ts.restaurants)


}
