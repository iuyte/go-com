package main

import (
  "time"

  "github.com/stianeikeland/go-rpio"
)

func main() {
  if err := rpio.Open(); err != nil {
    panic(err)
  }
  defer rpio.close

  pin := rpio.Pin(27)
  go func() {
    for {
      pin.Toggle()
      time.Sleep(time.Second)
    }
  }()
  time.Sleep(time.Minute)
}
