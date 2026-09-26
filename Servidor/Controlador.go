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
	  pr := s.CambiaEstado(mensaje, conn, nombre)
     	  if !pr {
	     s.UsuarioDesconectado(mensaje, conn, nombre)
	     for sala, _ := range s.salas {
	      	 if s.salas[sala].ContieneCliente(nombre) {
	      	    mensaje.Roomname = sala
	      	    s.salas[sala].DejaSala(mensaje, nombre, conn)
	      	 }
	     }
	     conn.Close()
	  }
     	  
     case "PUBLIC_TEXT":
     	  nombre := s.GetNombre(conn)
     	  s.MensajePublico(mensaje, conn, nombre)

     case "TEXT":
     	  nombre := s.GetNombre(conn)
	  s.TextoPrivado(mensaje, conn, nombre)
     
     case "USERS":
     	  s.ListaUsuarios(conn)
	  
     case "NEW_ROOM":
     	  nombre := s.GetNombre(conn)
     	  if mensaje.Roomname == "" {
	     s.UsuarioDesconectado(mensaje, conn, nombre)
	     for sala, _ := range s.salas {
	     	  if s.salas[sala].ContieneCliente(nombre) {
	      	     mensaje.Roomname = sala
	      	     s.salas[sala].DejaSala(mensaje, nombre, conn)
	      	  }
	     }
	     conn.Close()
	     return
	  }
     	  s.salas[mensaje.Roomname] = NuevaSala(mensaje.Roomname)
	  s.New_room(mensaje, nombre, conn)
	  s.salas[mensaje.Roomname].AgregaCliente(s.clientes[nombre])
	  
     case "INVITE":
     	  nombre := s.GetNombre(conn)
	  if len(mensaje.Usernames) == 0 {
	     s.UsuarioDesconectado(mensaje, conn, nombre)
	     for sala, _ := range s.salas {
	     	  if s.salas[sala].ContieneCliente(nombre) {
	      	     mensaje.Roomname = sala
	      	     s.salas[sala].DejaSala(mensaje, nombre, conn)
	      	  }
	     }
	     conn.Close()
	     return
	  }
	  sala, existe := s.salas[mensaje.Roomname]
	  if !existe {
	     s.NoExisteSala(mensaje, nombre, conn)
	     return
	  }
	  if !sala.ContieneCliente(nombre) {
	     s.FueraDeSala(mensaje, nombre, conn)
	     return
	  }
	  
	  usuarios := mensaje.Usernames
	  for i, cliente := range usuarios {
	      _, existe := s.clientes[cliente]
	      if !existe {
	      	 s.ClienteInexistente(mensaje, nombre, conn, cliente)
		 return
	      }
	      
	      if sala.ContieneCliente(cliente) || sala.invitados[cliente] {
	      	 usuarios = append(usuarios[:i], usuarios[i+1:]...)
	      }

	      sala.invitados[cliente] = true
	  }
	  
	  s.InvitaClientes(mensaje, nombre, usuarios)

     case "JOIN_ROOM":
     	  nombre := s.GetNombre(conn)
	  if mensaje.Roomname == "" {
	     s.UsuarioDesconectado(mensaje, conn, nombre)
	     for sala, _ := range s.salas {
	     	  if s.salas[sala].ContieneCliente(nombre) {
	      	     mensaje.Roomname = sala
	      	     s.salas[sala].DejaSala(mensaje, nombre, conn)
	      	  }
	     }
	     conn.Close()
	     return
	  }
	  
	  sala, existe := s.salas[mensaje.Roomname]
	  if !existe {
	     s.NoExisteSala(mensaje, nombre, conn)
	     return
	  }
	  if sala.ContieneCliente(nombre) {
	     return
	  }
	  if !sala.invitados[nombre] {
	     s.NoInvitado(mensaje, nombre, conn)
	     return
	  }
	  
	  s.UneCliente(mensaje, nombre, sala)
	  s.JoinedRoom(mensaje, nombre, conn, sala)


     case "ROOM_USERS":
     	  nombre := s.GetNombre(conn)
     	  sala, existe := s.salas[mensaje.Roomname]
	  if !existe {
	     s.NoExisteSala(mensaje, nombre, conn)
	     return
	  }
	  
	  if !sala.ContieneCliente(nombre) {
	     s.FueraDeSala(mensaje, nombre, conn)
	     return
	  }
     	  sala.ListaSala(conn)

     case "ROOM_TEXT":
     	  nombre := s.GetNombre(conn)
	  sala, existe := s.salas[mensaje.Roomname]
	  if !existe {
	     s.NoExisteSala(mensaje, nombre, conn)
	     return
	  }
	  if !sala.ContieneCliente(nombre) {
	     s.FueraDeSala(mensaje, nombre, conn)
	     return
	  }
	  sala.EnviaMensajeSala(mensaje, nombre, conn)

     case "LEAVE_ROOM":
     	  nombre := s.GetNombre(conn)
	  sala, existe := s.salas[mensaje.Roomname]
	  if !existe {
	     s.NoExisteSala(mensaje, nombre, conn)
	     return
	  }
	  if !sala.ContieneCliente(nombre) {
	     s.FueraDeSala(mensaje, nombre, conn)
	     return
	  }
	  sala.DejaSala(mensaje, nombre, conn)

     case "DISCONNECT":
     	  nombre := s.GetNombre(conn)
	  s.UsuarioDesconectado(mensaje, conn, nombre)
	  for sala, _ := range s.salas {
	      if s.salas[sala].ContieneCliente(nombre) {
	      	 mensaje.Roomname = sala
	      	 s.salas[sala].DejaSala(mensaje, nombre, conn)
	      }
	  }
	  conn.Close()

     default:
	  nombre := s.GetNombre(conn)
	  s.MsjInvalido(mensaje, nombre)
	  s.UsuarioDesconectado(mensaje, conn, nombre)
	  for sala, _ := range s.salas {
	      if s.salas[sala].ContieneCliente(nombre) {
	      	 mensaje.Roomname = sala
	      	 s.salas[sala].DejaSala(mensaje, nombre, conn)
	      }
	  }
	  conn.Close()
	  
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
