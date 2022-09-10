package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tarm/serial"
)

func getComPort() string {
	isport := false
	port := "COM"
	max := 20
	for i := 0; i < max; i++ {
		port = "COM" + strconv.Itoa(i)

		c := &serial.Config{Name: port, Baud: 9600}
		isport = true
		s, err := serial.OpenPort(c)
		if err != nil {
			isport = false
		}
		if isport {
			i = max
			s.Close()
		}
	}
	return port

}

func main() {
	fmt.Println("hello world    ....")
	port := getComPort()
	fmt.Printf("Port : %s\n", port)
	c := &serial.Config{Name: port, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("Unable to open port")
		log.Fatal(err)
	}
	line := ""
	buf := make([]byte, 1)
	on := true
	n := 0
	for on != false {
		n, err = s.Read(buf)
		if err != nil {
			fmt.Println("Error ")
			log.Fatal(err)
		}
		if string(buf[:n]) != "13" || string(buf[:n]) != "10" {
			line = line + string(buf[:n])
		}
		fmt.Printf("%s", buf[:n])
		//	cr := bytes.Equal(buf[:n], []byte(13))
		if string(buf[:n]) == "13" {
			fmt.Println(line)
			line = ""

		}
	}
}
