package main

import (
       "MyP/Comun"
       "net"
       "encoding/json"
       "fmt"
       "strings"
)

// Función que manda un mensaje privado a un cliente
func (s *Servidor) TextoPrivado(m comun.Mensaje, conn net.Conn, nombre string) {
     cliente, existe := s.clientes[m.Username]
     if !existe {
     	respuesta := comun.Mensaje{
     	       Type:      "RESPONSE",
	       Operation: "TEXT",
	       Result:    "NO_SUCH_USER",
	       Extra:     m.Username,
        }
        data, err := json.Marshal(respuesta)
        if err != nil {
     	   fmt.Printf("Error al recibir mensaje: %v\n", err)
        }
	 
        fmt.Printf(">>> %s\n", data)
        codificador := json.NewEncoder(conn)
        codificador.Encode(respuesta)
     	return
     }
     respuesta := comun.Mensaje{
     	       Type:     "TEXT_FROM",
	       Username: nombre,
	       Text:     m.Text,
     }
     data, err := json.Marshal(respuesta)
     if err != nil {
     	fmt.Printf("Error al recibir mensaje: %v\n", err)
     }
	 
     fmt.Printf(">>> %s\n", data)
     codificador := json.NewEncoder(cliente.conexion)
     codificador.Encode(respuesta)
     
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
     cliente.estado = m.Status
     for _, cliente := range s.clientes {
     	 if cliente.conexion == conn {
	    continue
	 }
	 respuesta := comun.Mensaje{
	      Type:     "NEW_STATUS",
	      Username: nombre,
	      Status:   m.Status,
	 }
	 data, err := json.Marshal(respuesta)
	 if err != nil {
	    fmt.Printf("Error al recibir mensaje: %v\n", err)
	 }
	 
	 fmt.Printf(">>> %s\n", data)
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
     data, err := json.Marshal(respuesta)
     if err != nil {
          fmt.Printf("Error al recibir mensaje: %v\n", err)
     }
	 
     fmt.Printf(">>> %s\n", data)
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
	  data, err := json.Marshal(respuesta)
 	  if err != nil {
	     fmt.Printf("Error al recibir mensaje: %v\n", err)
	  }
	 
	  fmt.Printf(">>> %s\n", data)
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
	 data, err := json.Marshal(respuesta)
	 if err != nil {
	    fmt.Printf("Error al recibir mensaje: %v\n", err)
	 }
	 
	 fmt.Printf(">>> %s\n", data)
	 
	 codificador := json.NewEncoder(cliente.conexion)
	 codificador.Encode(respuesta)
     }
}

// Función que manda el mensaje a todos
func (s *Servidor) UsuarioDesconectado(mensaje comun.Mensaje, conn net.Conn, nombre string) {
     for _, cliente := range s.clientes {
     	 if cliente.conexion == conn {
	    continue
	 }
	 respuesta := comun.Mensaje{
	      Type:     "DISCONNECTED",
	      Username: nombre,
	 }
	 data, err := json.Marshal(respuesta)
	 if err != nil {
	    fmt.Printf("Error al recibir mensaje: %v\n", err)
	 }
	 
	 fmt.Printf(">>> %s\n", data)
	 
	 codificador := json.NewEncoder(cliente.conexion)
	 codificador.Encode(respuesta)
     }
}

// Función que manda un mensaje después de intentarse crear una sala
func (s *Servidor) New_room(m comun.Mensaje, nombre string, conn net.Conn) {
     _, existe := s.salas[m.Roomname]
     if !existe {
     	respuesta := comun.Mensaje{
     	       Type:      "RESPONSE",
	       Operation: "NEW_ROOM",
	       Result:    "ROOM_ALREADY_EXISTS",
	       Extra:     m.Roomname,
        }
        data, err := json.Marshal(respuesta)
        if err != nil {
     	   fmt.Printf("Error al recibir mensaje: %v\n", err)
        }
	 
        fmt.Printf(">>> %s\n", data)
        codificador := json.NewEncoder(conn)
        codificador.Encode(respuesta)
     	return
     }
     respuesta := comun.Mensaje{
     	       Type:      "RESPONSE",
	       Operation: "NEW_ROOM",
	       Result:    "SUCCESS",
	       Extra:  	  m.Roomname,
     }
     data, err := json.Marshal(respuesta)
     if err != nil {
     	fmt.Printf("Error al recibir mensaje: %v\n", err)
     }
	 
     fmt.Printf(">>> %s\n", data)

     cliente, _ := s.clientes[nombre]
     codificador := json.NewEncoder(cliente.conexion)
     codificador.Encode(respuesta)
}


// Función que manda el mensaje de que no existe una sala
func (s *Servidor) NoExisteSala(m comun.Mensaje, nombre string, conn net.Conn) {
     
     respuesta := comun.Mensaje{
     	       Type:      "RESPONSE",
	       Operation: m.Type,
	       Result:    "NO_SUCH_ROOM",
	       Extra:  	  m.Roomname,
     }
     data, err := json.Marshal(respuesta)
     if err != nil {
     	fmt.Printf("Error al recibir mensaje: %v\n", err)
     }
	 
     fmt.Printf(">>> %s\n", data)

     cliente, _ := s.clientes[nombre]
     codificador := json.NewEncoder(cliente.conexion)
     codificador.Encode(respuesta)
}

// Función que manda un mensaje si un cliente fuera de una sala solicita la lista de dicha sala
func (s *Servidor) FueraDeSala(m comun.Mensaje, nombre string, conn net.Conn) {
     
     respuesta := comun.Mensaje{
     	       Type:      "RESPONSE",
	       Operation: m.Type,
	       Result:    "NOT_JOINED",
	       Extra:  	  m.Roomname,
     }
     data, err := json.Marshal(respuesta)
     if err != nil {
     	fmt.Printf("Error al recibir mensaje: %v\n", err)
     }
	 
     fmt.Printf(">>> %s\n", data)

     cliente, _ := s.clientes[nombre]
     codificador := json.NewEncoder(cliente.conexion)
     codificador.Encode(respuesta)
}

// Función que manda un mensaje si un cliente que es invitado a una sala no existe
func (s *Servidor) ClienteInexistente(m comun.Mensaje, nombre string, conn net.Conn, cliente string) {
     
     respuesta := comun.Mensaje{
     	       Type:      "RESPONSE",
	       Operation: "INVITE",
	       Result:    "NO_SUCH_USER",
	       Extra:  	  cliente,
     }
     data, err := json.Marshal(respuesta)
     if err != nil {
     	fmt.Printf("Error al recibir mensaje: %v\n", err)
     }
	 
     fmt.Printf(">>> %s\n", data)

     cl, _ := s.clientes[nombre]
     codificador := json.NewEncoder(cl.conexion)
     codificador.Encode(respuesta)
}

// Función que manda un mensaje si un cliente no fue invitado a una sala
func (s *Servidor) NoInvitado(m comun.Mensaje, nombre string, conn net.Conn) {
     
     respuesta := comun.Mensaje{
     	       Type:      "RESPONSE",
	       Operation: "JOIN_ROOM",
	       Result:    "NOT_INVITED",
	       Extra:  	  m.Roomname,
     }
     data, err := json.Marshal(respuesta)
     if err != nil {
     	fmt.Printf("Error al recibir mensaje: %v\n", err)
     }
	 
     fmt.Printf(">>> %s\n", data)

     cliente, _ := s.clientes[nombre]
     codificador := json.NewEncoder(cliente.conexion)
     codificador.Encode(respuesta)
}

// Función para anunciar la llegada de un nuevo cliente a una sala
func (s *Servidor) JoinedRoom(m comun.Mensaje, nombre string, conn net.Conn, sala *Sala) {
     
     for _, cliente := range sala.listaUsers {
     	 if cliente.conexion == conn {
	    continue
	 }
	 respuesta := comun.Mensaje{
	      Type:     "JOINED_ROOM",
	      Roomname: m.Roomname,
	      Username: nombre,
	 }
	 data, err := json.Marshal(respuesta)
	 if err != nil {
	    fmt.Printf("Error al recibir mensaje: %v\n", err)
	 }
	 
	 fmt.Printf(">>>prueba %s\n", data)
	 codificador := json.NewEncoder(cliente.conexion)
	 codificador.Encode(respuesta)
     }
}

// Función para cuando se reciba un mensaje inválido 
func (s *Servidor) MsjInvalido(m comun.Mensaje, conn net.Conn) {
     
     for _, cliente := range sala.listaUsers {
     	 if cliente.conexion == conn {
	    continue
	 }
	 respuesta := comun.Mensaje{
	      Type:      "RESPONSE",
	      Operation: m.Roomname,
	      Username: nombre,
	 }
	 data, err := json.Marshal(respuesta)
	 if err != nil {
	    fmt.Printf("Error al recibir mensaje: %v\n", err)
	 }
	 
	 fmt.Printf(">>>prueba %s\n", data)
	 codificador := json.NewEncoder(cliente.conexion)
	 codificador.Encode(respuesta)
     }
}
