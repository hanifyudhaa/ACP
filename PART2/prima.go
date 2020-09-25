package main

import (
  f "fmt"
  "time"
  )

func main(){

  var input int

  now := time.Now() 

  f.Print("Masukan Bilangan : ")
  f.Scan(&input)

  j:=0

  for i:=1; i<input; i++{
      if (input%i==0){
        j++
      }

  }

  if (j==1){
    f.Println("YA")
  } else {
    f.Println("TIDAK")
  }

    f.Println("time elapse:", time.Since(now)) 

}