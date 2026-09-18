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
     if s.clientes == nil {
     	s.clientes = make(map[string]ClienteConectado)
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

// Función que procesa un mensaje
func (s *Servidor) ProcesaMensaje(mensaje comun.Mensaje, conn net.Conn) {
     switch mensaje.Type {
     case "NEW_USER":
     	  s.NuevoUsuario(conn)

     case "STATUS":
     	  nombre := s.GetNombre(conn)
     	  s.CambiaEstado(mensaje, conn, nombre)
     	  
     case "PUBLIC_TEXT":
     	  nombre := s.GetNombre(conn)
     	  s.MensajePublico(mensaje, conn, nombre)

     case "TEXT":
     case "USERS":
     	  s.ListaUsuarios(conn)
	  
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

// Función para cambiar el estado de un usuario y avisar a los demás
func (s *Servidor) CambiaEstado(m comun.Mensaje, conn net.Conn, nombre string) {
     cliente, existe := s.clientes[nombre]
     if existe && cliente.estado == m.Status {
     	return
     }
     if m.Status != "ACTIVE" && m.Status != "AWAY" &&
     	m.Status != "BUSY" {
	return
     }
     for _, cliente := range s.clientes {
     	 if cliente.conexion == conn {
	    continue
	 }
	 respuesta := comun.Mensaje{
	      Type:     "NEW_STATUS",
	      Username: nombre,
	      Status:   m.Status,
	 }
	 codificador := json.NewEncoder(cliente.conexion)
	 codificador.Encode(respuesta)
     }
}



// Función que regresa la lista de usuarios de la sala principal
func (s *Servidor) ListaUsuarios(conn net.Conn) {
     users := make(map[string]string)
     for nombre, cliente := range s.clientes {
	      users[nombre] = cliente.estado
	  }
	  respuesta := comun.Mensaje{
	  	    Type:  "USER_LIST",
		    Users: users,
	  }
	  codificado := json.NewEncoder(conn)
	  codificado.Encode(respuesta)
}

// Función que avisa a los demás usuarios (si hay) que llegó alguien nuevo
func (s *Servidor) NuevoUsuario(conn net.Conn) {
     
     nombre := s.GetNombre(conn)
     for _, cliente := range s.clientes {
	      
	  if cliente.conexion == conn {
	     continue
	  }
	  codificado := json.NewEncoder(cliente.conexion)
	  respuesta := comun.Mensaje{
	  	    Type:     "NEW_USER",
		    Username: nombre,
	  }
	  codificado.Encode(respuesta)
     }
}

// Función que manda el mensaje a todos
func (s *Servidor) MensajePublico(mensaje comun.Mensaje, conn net.Conn, nombre string) {
     msj := strings.TrimSpace(mensaje.Text)
     if len(msj) == 0 {
     	return
     }
     for _, cliente := range s.clientes {
     	 if cliente.conexion == conn {
	    continue
	 }
	 respuesta := comun.Mensaje{
	      Type:     "PUBLIC_TEXT_FROM",
	      Username: nombre,
	      Text:     mensaje.Text,
	 }
	 codificador := json.NewEncoder(cliente.conexion)
	 codificador.Encode(respuesta)
     }
}

// Función auxiliar que regresa el nombre del cliente mediante la conexión
func (s *Servidor) GetNombre(conn net.Conn) string {
     for nombre, cliente := range s.clientes {
     	 if cliente.conexion == conn {
	    return nombre
	 }
     }
     return ""
}