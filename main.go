package main

import (
	"fmt"
)

func main() {

  // Hello World  - L1
  fmt.Println("Hello World !");

  // Hello World - L2
  fmt.Println(fmt.Sprintf("Hello %s !", "World"));

  //Hello World - L3
  var a string = "Hello"
  var b string = "World"
  fmt.Println(a, b, "!");

  //Hello World - L4
  m := "Hello"
  n := "World"
  fmt.Println(m , n, "!")

  //Hello World - L4
  p := &m
  q := &n
  fmt.Println(*p , *q, "!")

  // Hello World - L5
  x := []rune{'H', 'e', 'l', 'l', 'o'}
  y := []rune{'W', 'o', 'r', 'l', 'd'}
  fmt.Println(string(x), string(y), "!");

  // Hello World - L6
  s := func() string { return "Hello"}
  t := func() string { return "World"}
  fmt.Println(s(), t(), "!");

  g := func(fn1, fn2 func() string) string {
    return fmt.Sprintf("%s %s", fn1(), fn2())
  }
  fmt.Println(g(s, t), "!");
}