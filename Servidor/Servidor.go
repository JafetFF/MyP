package main

import (
       "fmt"
       "log"
       "net"
)

type Servidor struct {
     puerto   string
}

// Función que regresa un servidor
func NuevoServidor(puerto string) *Servidor {
     return &Servidor{
     	    puerto: puerto,
     }
}

// Función que inicia el servidor dado
func (s *Servidor) Iniciar() {
      listener, err := net.Listen("tcp", s.puerto)
      if err != nil {
      	 log.Fatalf("Error")
      }
      defer listener.Close()
      fmt.Printf("Escuchando en el puerto %s\n", s.puerto)

      for {
      	  conn, err := listener.Accept()
	  if err != nil {
	     log.Printf("Error al aceptar conexión")
	     continue
	  }
	  go s.AtiendeConexion(conn)
      }
}

// Función que acepta conexiones
func (s *Servidor) AtiendeConexion(conn net.Conn) {
     defer conn.Close()
     fmt.Printf("New Conexion desde %s\n", conn.RemoteAddr().String())
     conn.Write([]byte("Mensaje recibido correctamente"))
}