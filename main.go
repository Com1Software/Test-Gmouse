package main

import (
	"fmt"
      "github.com/tarm/serial"
      "log"
)

func main() {
	fmt.Println("hello world    ....")
      c := &serial.Config{Name: "COM7", Baud: 115200}
      s, err := serial.OpenPort(c)
      if err != nil {
              log.Fatal(err)
      }

      n, err := s.Write([]byte("test"))
      if err != nil {
              fmt.Println("Test0")
              log.Fatal(err)
      }

      buf := make([]byte, 128)
      n, err = s.Read(buf)
      if err != nil {
               fmt.Println("Test1")
              log.Fatal(err)
      }
       fmt.Println("Start Test...")
      log.Print("%q", buf[:n])
      fmt.Println("... End Test")
}
