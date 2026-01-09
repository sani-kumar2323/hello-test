package main

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"net/http"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCP Server Running on :9000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		log.Println("Received:", msg)

		// send to db microservice
		http.Post("http://localhost:8082/save", "application/json",
			bytes.NewBuffer([]byte(`{"meter_id":"M1","value":`+msg+`}`)),
		)
	}
}

