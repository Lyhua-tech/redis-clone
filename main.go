package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	// making a server for running with port 6379
	fmt.Println("Listening on port :6379")
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}


	// recieve the request.
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	aof, err := NewAof("database.aof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer aof.Close()

	aof.Read(func(value Value) {
		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			return
		}

		handler(args)
	})

	// Run the loop for recieving command like 
	// SET, GET, HSET ,HGET
	for {

		// prepare to read the network as a chunk not one by one.
		resp := NewResp(conn)
		value, err := resp.Readed()
		if err != nil {
			fmt.Println(err)
			return
		}
		if value.typ != "array" {
			fmt.Println("Invalid request, expected an array")
			continue
		}

		if len(value.array) == 0{
			fmt.Println("Invalid request, expected an array length > 0")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		writer := NewWriter(conn)

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid request:", command)
			writer.Write(Value{typ: "string", str: ""})
			continue
		}

		if command == "SET" || command == "HSET"{
			aof.Write(value)
		}

		result := handler(args)
		writer.Write(result)
	}
}