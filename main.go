package main

import (
  "fmt"
)

func pi() {
  fmt.Printf("init\n")
}
func pm() {
  fmt.Printf("main\n")
}

func init() {
  pi()
}

func main() {
  pm()
}
