package main

import (
       "fmt"
       "encoding/json"
       "log"
       "net"
)

type Servidor struct {
     puerto string
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
      	 log.Fatalf("Error: %v", err)
      }
      defer listener.Close()
      fmt.Printf("Escuchando en el puerto %s\n", s.puerto)

      for {
      	  conn, err := listener.Accept()
	  if err != nil {
	     log.Println("Error al aceptar conexión")
	     continue
	  }
	  go s.AtiendeConexion(conn)
      }
}

// Función que atiende conexiones
func (s *Servidor) AtiendeConexion(conn net.Conn) {
     defer conn.Close()
     fmt.Printf("Nueva Conexión desde %s\n", conn.RemoteAddr().String())

     decodificado := json.NewDecoder(conn)
     var m mensaje
     err := decodificado.Decode(&m)
     if err != nil {
     	respuesta := mensaje{
	Type: 	   "RESPONSE",
     	Operation: "INVALID",
	Result:    "INVALID",
	}
	codificado := json.NewEncoder(conn)
	err = codificado.Encode(respuesta)
	if err != nil {
	   fmt.Printf("Error al enviar respuesta %v\n", err)
	}
	return
     }
     if m.Type != "IDENTIFY" {
     	respuesta := mensaje{
	Type: 	   "RESPONSE",
     	Operation: "INVALID",
	Result:    "NOT_IDENTIFIED",
	}
	codificado := json.NewEncoder(conn)
	err = codificado.Encode(respuesta)
	if err != nil {
	   fmt.Printf("Error al enviar respuesta %v\n", err)
	}
	return
     }
     conn.Write([]byte("pasó"))
	    
}