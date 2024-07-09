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
   // different modelling of decay factor can give different results.
  var history_choices []int = []int{1,1,2,3,2,2,0,0,2,2}
  var feedback []int = []int{1,1,0,0,1,1,0,0,0,0}
  var tau float64 = 1 // accumulated positive reward weight 
  var phi float64 =1 // accumulated negative reward weight 
  var positive_reward_weight float64 = 1
  var negative_reward_weight float64 = 1

  


  var ts *ThompsonSampling = InitThompsonSampling(num_restaurants,tau,phi,positive_reward_weight,negative_reward_weight)

  
  // run trials
  for trial:=0;trial< len(history_choices) + num_trials;trial++ {
    if(trial < len(history_choices)){
      if(history_choices[trial]>=num_restaurants){
        panic("Invalid input history_choices: restaurant not within range")
      }
      ts.Feedback(history_choices[trial],feedback[trial])
      fmt.Println("values of params: ")
      fmt.Println(ts.restaurants)
    } else {
      fmt.Println("Trial ",trial-len(history_choices))
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
        ts.Feedback(suggested_restaurant,1)
      } else {
        ts.Feedback(suggested_restaurant,0)
      }
    }

  }
  fmt.Println("Updated values of params: ")
  fmt.Println(ts.restaurants)


}
