package main

import (
       "fmt"
       "log"
       "net"
)

func main() {
      listener, err := net.Listen("tcp", ":8080")
      if err != nil {
      	 log.Fatalf("Error")
      }
      defer listener.Close()
      fmt.Printf("Escuchando en el puerto 8080...")

      for {
      	  conn, err := listener.Accept()
	  if err != nil {
	     log.Printf("Error al aceptar conexión")
	     continue
	  }
	  go conexion(conn)
      }
}

func conexion(conn net.Conn) {
     defer conn.Close()
     fmt.Printf("New Conexion desde", conn.RemoteAddr().String())
     conn.Write([]byte("Mensaje recibido correctamente"))
}