package main

import (
	"bufio"
	"fmt"
	"net"
)

// go build -o net1.exe . && ./net1.exe

func main1() {

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("")
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic("")
		}
		go handleconn(conn)
	}
}

func handleconn(con net.Conn) {

	name := con.RemoteAddr().String()

	fmt.Println("conect", name)

	con.Write([]byte("Hello, " + name + "\n\r"))

	defer con.Close()

	scanner := bufio.NewScanner(con)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "Exit" {
			con.Write([]byte("Bay" + name))
			fmt.Println("disconnect", name)
			return
		} else {
			fmt.Println(text + " " + name)
			con.Write([]byte("Your Enter " + text))
		}
	}
}
