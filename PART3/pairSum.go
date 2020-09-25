package main

import f "fmt"

func main(){

  pairSum([]int{1,2,3,4,5,6}, 6)
  pairSum([]int{2,5,9,11}, 11)

}

func pairSum(arr []int , target int){
  hit :=0

  if len(arr)%2==0 {
    hit=len(arr)/2
  } else{
    hit=(len(arr)+1)/2
  }

  for i:=0; i<hit; i++{
    for j:=len(arr)-1; j>=hit; j--{
      if arr[i] + arr[j] == target && arr[i]!= arr[j]{
        f.Println(i,j)
        f.Println(arr[i],arr[j])
      }
    }
  }
}