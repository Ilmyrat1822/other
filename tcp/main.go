package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			panic(err)

		}
		/*io.WriteString(con, "Hello from server\n")
		fmt.Fprintln(con, "Good good")
		con.Close()*/
		go handle(con)

	}
}

func handle(con net.Conn) {

	scanner := bufio.NewScanner(con)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(con, " You said: %s\n", line)
	}
	defer con.Close()

	fmt.Println("Code never got here")
}
