package main

import (
       "fmt"
       "encoding/json"
       "log"
       "net"
       "MyP/Comun"
       "strings"
)

type Servidor struct {
     puerto   string
     clientes map[string]*ClienteConectado
}

type ClienteConectado struct {
     nombre   string
     estado   string
     conexion net.Conn
}

// Función que regresa un servidor
func NuevoServidor(puerto string) *Servidor {
     return &Servidor{
     	    puerto:   puerto,
	    clientes: make(map[string]*ClienteConectado),
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
     if s.clientes == nil {
     	s.clientes = make(map[string]*ClienteConectado)
     }
     
     fmt.Printf("Nueva Conexión desde %s\n", conn.RemoteAddr().String())

     decodificado := json.NewDecoder(conn)
     var m comun.Mensaje
     err := decodificado.Decode(&m)
     if err != nil {
     	respuesta := comun.Mensaje{
	Type: 	   "RESPONSE",
     	Operation: "INVALID",
	Result:    "INVALID",
	}
	codificado := json.NewEncoder(conn)
	err = codificado.Encode(respuesta)
	if err != nil {
	   fmt.Printf("Error al enviar respuesta %v\n", err)
	}
	fmt.Printf(">>> Mensaje del servidor: %s\n", respuesta)
	return
     }
     if m.Type != "IDENTIFY" {
     	respuesta := comun.Mensaje{
	Type: 	   "RESPONSE",
     	Operation: "INVALID",
	Result:    "NOT_IDENTIFIED",
	}
	codificado := json.NewEncoder(conn)
	err = codificado.Encode(respuesta)
	if err != nil {
	   fmt.Printf("Error al enviar respuesta %v\n", err)
	}
	fmt.Printf("Mensaje del servidor: %s\n", respuesta)
	return
     }
     if strings.TrimSpace(m.Username) == "" {
     	respuesta := comun.Mensaje{
	Type: 	   "RESPONSE",
     	Operation: "INVALID",
	Result:    "INVALID",
	}
	codificado := json.NewEncoder(conn)
	codificado.Encode(respuesta)
	data, err := json.Marshal(respuesta)
	if err != nil {
	   fmt.Printf("Error al recibir mensaje: %v\n", err)
	}
	 
	fmt.Printf(">>> %s\n", data)
	
	return
     }

     msj, err := json.Marshal(m)
	 if err != nil {
	    fmt.Printf("Error al recibir mensaje: %v\n", err)
	 }

     fmt.Printf("<<< %s\n", msj)

     if _, existe :=
     s.clientes[m.Username]; existe{
     	respuesta := comun.Mensaje{
		  Type:      "RESPONSE",
		  Operation: "IDENTIFY",
		  Result:    "USER_ALREADY_EXISTS",
		  Extra:     m.Username,
	}
	codificado := json.NewEncoder(conn)
	codificado.Encode(respuesta)
	data, err := json.Marshal(respuesta)
	 if err != nil {
	    fmt.Printf("Error al recibir mensaje: %v\n", err)
	 }
	 
	 fmt.Printf(">>> %s\n", data)
	return
     }

     s.clientes[m.Username] = &ClienteConectado{
     	nombre:   m.Username,
	estado:   "ACTIVE",
	conexion: conn,
     }

     defer func() {
     	   delete(s.clientes, m.Username)
     	   conn.Close()
     }()

     respuesta := comun.Mensaje{
     Type:         "RESPONSE",
     Operation:	   "IDENTIFY",
     Result:  	   "SUCCESS",
     Extra:   	   m.Username,
     }

     data, err := json.Marshal(respuesta)
     if err != nil {
         fmt.Printf("Error al recibir mensaje: %v\n", err)
     }
	 
     fmt.Printf(">>> %s\n", data)
     
     codificado := json.NewEncoder(conn)
     
     err = codificado.Encode(respuesta)
     if err != nil {
     	fmt.Printf("Error al enviar respuesta %v\n", err)
     }

     s.NuevoUsuario(conn)

     for {
     	 
     	 var mensaje comun.Mensaje

	 err := decodificado.Decode(&mensaje)
	 if err != nil {
	    fmt.Printf("Cliente desconectado: %v\n", err)
	    return
	 }
	 data, err := json.Marshal(mensaje)
	 if err != nil {
	    fmt.Printf("Error al recibir mensaje: %v\n", err)
	 }
	 
	 fmt.Printf("<<< %s\n", data)
	 
	 
	 s.ProcesaMensaje(mensaje, conn)

     } 
}



