package main

import (
       "fmt"
       "log"
       "net"
)

func main() {
      PORT := ":8080"

      listener, err := net.Listen("tcp", port)
      if err != nil {
      	 log.Fatalf("Error")
      }
      defer listener.Close()
      fmt.Printf("Escuchando")

      for {
      	  conn, err := listener.Accept()
	  if err != nil {
	     log.Printf("Error 2")
	     continue
	  }
	  go handleConnection(conn)
}

func handleConnection(conn net.Conn) {
     defer conn.Close()
     fmt.Printf("New Conexion desde", conn.RemoteAddr().String())
     conn.Write([]byte("Hola"))
}