package main

import f "fmt"

func main(){

  var x, y int

  f.Print("Masukan X : ")
  f.Scan(&x)
  f.Print("Masukan N : ")
  f.Scan(&y)
  temp := alas

  for i:= 1; i<y; i++{
    temp=x*temp
    f.Println(temp)
  }

  f.Println("X^N adalah =  ", temp)

}