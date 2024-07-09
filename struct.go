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
  tau float64; //prev_accumulated_positive_weight 
  phi float64; //prev_accumulated_negative_weight
  positive_reward_weight float64;
  negative_reward_weight float64;
}


func InitThompsonSampling(num_restaurants int,tau,phi,positive_reward_weight,negative_reward_weight float64) (*ThompsonSampling) {
  var restaurants []Params = make([]Params,num_restaurants)
  for i:=0;i<num_restaurants;i++ {
    restaurants[i] = Params{alpha:1,beta:1}
  }
  
  return &ThompsonSampling{num_restaurants:num_restaurants,restaurants:restaurants, tau:tau, phi:phi,positive_reward_weight:positive_reward_weight, negative_reward_weight:negative_reward_weight}
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

func (ts *ThompsonSampling) Feedback(i int, feedback int) {
  var alpha = ts.restaurants[i].alpha
  var beta = ts.restaurants[i].beta 
  if(feedback==0){
    ts.restaurants[i].beta = ts.phi * (beta)+ ts.negative_reward_weight*1 // here ri is 1
  } else {
    ts.restaurants[i].alpha = ts.tau * (alpha) + ts.positive_reward_weight*1
  }
    
}
