package main

import (
       "fmt"
       "encoding/json"
       "log"
       "net"
       "MyP/Comun"
)

type Servidor struct {
     puerto   string
     clientes map[string]ClienteConectado
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
	    clientes: make(map[string]ClienteConectado),
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
	return
     }

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
	return
     }

     s.clientes[m.Username] = ClienteConectado{
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
     
     codificado := json.NewEncoder(conn)
     err = codificado.Encode(respuesta)
     if err != nil {
     	fmt.Printf("Error al enviar respuesta %v\n", err)
     }

     for {
     	 var mensaje comun.Mensaje

	 err := decodificado.Decode(&mensaje)
	 if err != nil {
	    fmt.Printf("Cliente desconectado: %v\n", err)
	    return
	 }

	 fmt.Printf("Mensaje recibido: %+v\n", mensaje)

	 s.ProcesaMensaje(mensaje, conn)

     } 
}

// Función que procesa un mensaje
func (s *Servidor) ProcesaMensaje(mensaje comun.Mensaje, conn net.Conn) {
     switch mensaje.Type {
     case "PUBLIC_TEXT":
     case "TEXT":
     case "USERS":
     	  usuarios := make(map[string]string)

	  for nombre, cliente := range s.clientes {
	      usuarios[nombre] = cliente.estado
	  }
	  respuesta := comun.Mensaje{
	  	    Type:  "USER_LIST",
		    Users: usuarios,
	  }
	  codificado := json.NewEncoder(conn)
	  codificado.Encode(respuesta)
	  
     case "NEW_ROOM":
     case "INVITE":
     case "JOIN_ROOM":
     case "ROOM_USERS":
     case "ROOM_TEXT":
     case "LEAVE_ROOM":
     case "DISCONNECT":
     default:
     // cuando el mensaje sea inválido
     }
}