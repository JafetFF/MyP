package main

import (
       "MyP/Comun"
       "net"
       "unicode/utf8"
)

// Función que procesa un mensaje
func (s *Servidor) ProcesaMensaje(mensaje comun.Mensaje, conn net.Conn, nombre string) {
     switch mensaje.Type {
     case "STATUS":	  
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
     	  
     	  s.MensajePublico(mensaje, conn, nombre)

     case "TEXT":
     	  
	  if mensaje.Text == "" {
	     s.MsjInvalido(mensaje, nombre)
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
	  s.TextoPrivado(mensaje, conn, nombre)
     
     case "USERS":
     	  s.ListaUsuarios(conn)
	  
     case "NEW_ROOM":
     	  
     	  if mensaje.Roomname == "" {
	     s.MsjInvalido(mensaje, nombre)
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
	  if utf8.RuneCountInString(mensaje.Roomname) > 16 {
	     motivo := "Las salas deben tener a lo sumo 16 caracteres"
	     mensaje.Text = motivo
	     s.MsjInvalido(mensaje, nombre)
	     return
	  }
     	  
	  s.New_room(mensaje, nombre, conn)
	  
	  s.salas[mensaje.Roomname].AgregaCliente(s.clientes[nombre])
	  
     case "INVITE":
     	  
	  if len(mensaje.Usernames) == 0 {
	     s.MsjInvalido(mensaje, nombre)
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
	  
	  usuarios := []string{}
	  
	  for _, cliente := range mensaje.Usernames {
	      _, existe := s.clientes[cliente]
	      if !existe {
	      	 s.ClienteInexistente(mensaje, nombre, conn, cliente)
		 return
	      }
	      
	      if !sala.ContieneCliente(cliente) && !sala.invitados[cliente] {
	      	 usuarios = append(usuarios, cliente)
		 sala.invitados[cliente] = true
	      } 	      
	  }
	  
	  s.InvitaClientes(mensaje, nombre, usuarios)

     case "JOIN_ROOM":
     	  
	  if mensaje.Roomname == "" {
	     s.MsjInvalido(mensaje, nombre)
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
     	  
     	  if mensaje.Roomname == "" {
	     s.MsjInvalido(mensaje, nombre)
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
     	  sala.ListaSala(conn)

     case "ROOM_TEXT":
     	  
	  if mensaje.Text == "" {
	     s.MsjInvalido(mensaje, nombre)
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
	  sala.EnviaMensajeSala(mensaje, nombre, conn)

     case "LEAVE_ROOM":
     	  
	  if mensaje.Roomname == "" {
	     s.MsjInvalido(mensaje, nombre)
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
	  sala.DejaSala(mensaje, nombre, conn)
	  
	  if sala.EsVacia() {
	     delete(s.salas, mensaje.Roomname)
	  }

     case "DISCONNECT":
     	  
	  s.UsuarioDesconectado(mensaje, conn, nombre)
	  for sala, _ := range s.salas {
	      if s.salas[sala].ContieneCliente(nombre) {
	      	 mensaje.Roomname = sala
	      	 s.salas[sala].DejaSala(mensaje, nombre, conn)
	      }
	  }
	  conn.Close()

     default:
	  
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
