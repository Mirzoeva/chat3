package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func send_messages(conn net.Conn){
	var strEcho string
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	strEcho = in.Text()
	fmt.Println(strEcho)
	if n, err := conn.Write([]byte(strEcho)); n == 0 || err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
		return
	}
}

func main() {
	servAddr := "lab.posevin.com:10059"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Print("Username: ")
	var log string
	_, err = fmt.Scanln(&log)
	log += "_"
	if err != nil {
		fmt.Println("Error:", err)
	}
	if n, err := conn.Write([]byte(log)); n == 0 || err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
		return
	}
	for {
		go send_messages(conn)
		reply := make([]byte, 1024)
		n, err := conn.Read(reply)
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}
		println(string(reply[0:n]))
	}

}

