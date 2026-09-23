package main

import (
       "MyP/Comun"
       "net"
)

// Función que procesa un mensaje
func (s *Servidor) ProcesaMensaje(mensaje comun.Mensaje, conn net.Conn) {
     switch mensaje.Type {
     case "STATUS":
     	  nombre := s.GetNombre(conn)
     	  s.CambiaEstado(mensaje, conn, nombre)
     	  
     case "PUBLIC_TEXT":
     	  nombre := s.GetNombre(conn)
     	  s.MensajePublico(mensaje, conn, nombre)

     case "TEXT":
     	  nombre := s.GetNombre(conn)
	  s.TextoPrivado(mensaje, conn, nombre)
     
     case "USERS":
     	  s.ListaUsuarios(conn)
	  
     case "NEW_ROOM":
     case "INVITE":
     case "JOIN_ROOM":
     case "ROOM_USERS":
     case "ROOM_TEXT":
     case "LEAVE_ROOM":
     case "DISCONNECT":
     	  nombre := s.GetNombre(conn)
     	  s.Desconectado(mensaje, conn, nombre)

     default:
     // cuando el mensaje sea inválido
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