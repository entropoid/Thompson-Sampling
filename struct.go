package main

import (
  "gonum.org/v1/gonum/stat/distuv"
  "golang.org/x/exp/rand"
  "time"
)
var src = rand.NewSource(uint64(time.Now().UnixNano()))
var rng = rand.New(src)

type Params struct{
  alpha float64;
  beta float64;
}

func (p Params) Sample() (float64) {
  beta_distribution := distuv.Beta{Alpha:p.alpha, Beta:p.beta, Src:rng}
  var sampled_value float64 = beta_distribution.Rand()
  return sampled_value
}


type ThompsonSampling struct {
  num_restaurants int;
  restaurants []Params;
  decay_factor float64;
}


func ContextualThompsonSampling(history_choices []int, feedback []int, num_restaurants int, decay_factor float64) (*ThompsonSampling) {
  if(len(history_choices)!=len(feedback)){
    panic("len of history_choices and feedback should be same")
  }
  var restaurants []Params = make([]Params,num_restaurants)
  for i:=0;i<num_restaurants;i++ {
    restaurants[i] = Params{alpha:1,beta:1}
  }
  for i:=0;i<len(history_choices);i++ {
    for j,_ := range(restaurants){
      restaurants[j].alpha *= decay_factor
      restaurants[j].beta *= decay_factor
    }
    var restaurant_i = history_choices[i]
    var ri int= feedback[i]
    if(ri==1) {
      restaurants[restaurant_i].alpha += 1
    } else {
      restaurants[restaurant_i].beta += 1
    }
  }
  
  return &ThompsonSampling{num_restaurants:num_restaurants,restaurants:restaurants, decay_factor:decay_factor}

}


func (ts *ThompsonSampling) Sample() ([]float64) {
  // returns list of random values obtained by sampling
  var samples []float64 = make([]float64,ts.num_restaurants)
  for i, restaurant := range ts.restaurants {
    samples[i] = restaurant.Sample()
  }
  return samples 
}

func (ts *ThompsonSampling) Choose( samples []float64) int {
  // returns index with highest values from samples 
  var suggested_restaurant = -1
  var max_val float64 = -1
  for i,val := range(samples){
    if(val>max_val){
      suggested_restaurant = i
      max_val = val 
    }
  }
  return suggested_restaurant
}

func (ts *ThompsonSampling) Feedback(i int, feedback bool) {
  for j,_ := range(ts.restaurants){
      ts.restaurants[j].alpha *= ts.decay_factor
      ts.restaurants[j].beta *= ts.decay_factor
    }

  if(!feedback){
    ts.restaurants[i].beta += 1
  } else {
    ts.restaurants[i].alpha += 1
  }
    
}
