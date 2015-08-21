package main

import (
  "github.com/m0dd3r/reddit"
  "log"
  "fmt"
)

func main() {
  items, err := reddit.Get("golang")
  if err != nil {
    log.Fatal(err)
  }
  for _, item := range items{
    fmt.Println(item)
  }
}
