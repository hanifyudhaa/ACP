package main

import f "fmt"

func main(){

  var input int

  f.Print("Masukan Angka : ")
  f.Scan(&input)

  for i:=input; i>0; i--{
    for j:=1; j<=i/2; j++{
      f.Print(" ")
    }
    for k:=input; k>=i; k--{
      f.Print("*")
    }
    f.Println("")
  }

}