package main

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

func main() {
	fmt.Println("hello world    ....")
	c := &serial.Config{Name: "COM5", Baud: 9600}
	fmt.Println("testa")
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("Testb")
		log.Fatal(err)
	}

	n, err := s.Write([]byte("test"))
	if err != nil {
		fmt.Println("Test0")
		log.Fatal(err)
	}

	buf := make([]byte, 1280)
	n, err = s.Read(buf)
	if err != nil {
		fmt.Println("Test1")
		log.Fatal(err)
	}
	fmt.Println("Start Test...")
	log.Print("%s", string(buf[:n]))
	fmt.Println("... End Test")
}
