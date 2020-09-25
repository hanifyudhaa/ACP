package main

import (
  f "fmt"
  "sort"
)

func main (){

  meanMedian( []float64 {3,2,1,5,4} )

}

func meanMedian ( arr[] float64){
  var total float64 = 0

  for _, sum := range arr{
    total += sum
  }

  f.Println("Mean = ", total/float64(len(arr)))


  sort.Float64s(arr)
  f.Println(arr)

  tengah := len(arr)/2

  if len(arr) % 2 == 0 {
    f.Println(tengah)
    f.Println("Median = ", (arr[tengah]+arr[tengah-1])/2)

  } else {
    f.Println(tengah)
    f.Println("Median = ", arr[tengah])

  }


}