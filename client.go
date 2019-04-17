package main

import (
	"fmt"
	"net"
	"os"
)
var connections [1000] net.Conn;

func differentiate(s []byte) (a, b string){
	n:=0;
	for i:=0; i< len(s); i++{
		if s[i] == '_'{
			n=i;
			break
		}
	}
	return string(s[0:n]), string(s[n+1 : len(s)])
}

func write_to_clients(s string){
	for _, i := range connections {
		if i!=nil {
			i.Write([]byte(s))
		} else {
			break
		}
	}
}
func main() {
	dictionary:= make(map[string]string);
	l, err := net.Listen("tcp", "lab.posevin.com:10059")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	length_conn:=0
	for {
		conn, err := l.Accept()
		connections[length_conn] = conn
		length_conn++
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn, dictionary)
	}
}

func handleRequest(conn net.Conn, dictionary map [string]string ) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
	}
	login, password := differentiate(buf)
	if _, ok := dictionary[login]; !ok {
		dictionary[login] = password;
	}
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		s:=login +" : " + string(buf[0:n]);
		write_to_clients(s);
	}
	conn.Close()
}