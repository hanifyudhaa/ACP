package main

import f "fmt"

func main(){

  var bilangan int

  f.Print("Masukan Angka : ")
  f.Scan(&bilangan)

  for i:=1; i<=bilangan; i++{
    if bilangan % i == 0{
      f.Println(i)
    }
  }

}