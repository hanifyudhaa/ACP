package main

import f "fmt"

func main(){

  var input, result string

  f.Print("Masukan Kata : ")
  f.Scan(&input)

  for _, v := range input {
      result = string(v) + result
      f.Println(string(v))
      f.Println(result)
   }

  if input==result {
    f.Println("Palindrome")
  }else{
    f.Println("Bukan Palindrome")
  }

}