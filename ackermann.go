package main

import "fmt"

func ackermann(m uint, n uint) uint {
  if m == 0 {
    return  n + 1 
  }
  if n == 0 {
    return ackermann(m - 1, 1)
  }
  return ack(m - 1, ackermann(m, n - 1))
}

func main() {
  fmt.Printf("%d", ackermann(4, 2))
}
