package main

import f "fmt"

func main(){

  var input int

  f.Print("Masukan Bilangan : ")
  f.Scan(&input)

  j:=0

  for i:=1; i<input; i++{
      if (input%i==0){
        j++
      }

  }

  if (j==3){
    f.Println("YA")
  } else {
    f.Println("TIDAK")
  }

}