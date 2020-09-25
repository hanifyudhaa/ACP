package main

import (
  f "fmt"

)
func tugas1(studentScore int){
  if (studentScore >= 80 && studentScore<=100){
    f.Println("A")
  }else if(studentScore>=65 && studentScore <=79){
    f.Println("B")
  }else if(studentScore>=50 && studentScore <=64){
    f.Println("C")
  }else if(studentScore>=35 && studentScore <=49){
    f.Println("D")
  }else if(studentScore>=0 && studentScore <=34){
    f.Println("E")
  }else{
    f.Println("Invalid")
  }
}

func tugas2(height int){
  for i := 1; i <= height*height; i++ {
    if (i%2 != 0 && i%3 !=0){
      f.Print("@")
     }else if (i%2 == 0 && i%3 !=0){
      f.Print("$")
     }else{
      f.Print("*")
     }
     
     if (i%height==0){
      f.Println("")
    }
  } 
}

func tugas3(x int){
  for i := 1; i <= x; i++ {
    for j := 1; j <= x; j++ {
      f.Print(i*j,"\t")
     }
     f.Println("")
  }

}

func tugas4(x string){

}

func tugas5(x string){

  // for _, r := range x {
  //       c := string(r)
  //       f.Println(c)
  // }

  // newArr:= strings.Split(x, "")
  // f.Println(newArr[2])

  for j := 0; j < len(x); j++ {
  
    a := string(x[j])
    temp:=0

    for i := 0; i < len(x); i++ {
      b := string(x[i])
      if (a!=b){
        temp++
      }
    }

     if (temp == len(x)-1){
       f.Print(a,"\t")
     }
  }

}

func main() {
  
  var a,b,c int 
  var e string

  f.Println("\nTugas 1 Grade Nilai \n-------------------\n")
  f.Print("Masukkan Nilai : ")
  f.Scan(&a)
  tugas1(a)

  f.Println("\nTugas 2 Theree Cols Box \n-------------------\n")
  f.Print("Masukkan Nilai : ")
  f.Scan(&b)
  tugas2(b)

  f.Println("\nTugas 3 Cetak Tabel Perkalian \n-------------------\n")
  f.Print("Masukkan Nilai : ")
  f.Scan(&c)
  tugas3(c)

  f.Println("\nTugas 5 Angka Muncul Sekali \n-------------------\n")
  f.Print("Masukkan Angka : ")
  f.Scan(&e)
  tugas5(e)
}