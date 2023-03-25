package main

import (
  "fmt"
)

func pi() {
  fmt.Prinf("init\n")
}
func pm() {
  fmt.Prinf("main\n")
}

func init() {
  pi()
}

func main() {
  pm()
}
