package main

import (
       "strings"
       "encoding/json"
       "MyP/Comun"
)

// Función auxiliar que interpreta lo que el cliente solicita al servidor
func InterpretaMensaje(entrada string, codificado *json.Encoder) {
     texto := strings.TrimSpace(entrada)
     palabra := strings.Fields(texto)
     if len(palabra) == 0 {
     	return
     }
     switch palabra[0] {
     case "USERS": // lista de usuarios
     	  codificado.Encode(comun.Mensaje{
		Type: "USERS",
	  })
     case "STATUS": // estado
     	  if len(palabra) != 2 {
	     return
	  }
     	  codificado.Encode(comun.Mensaje{
		Type:   "STATUS",
		Status: palabra[1],
	  })
     case "TEXT": // msj privado
     	  if len(palabra) < 3 { // desconectaremos depués al usuario
	     
	     return
	  } 
     	  codificado.Encode(comun.Mensaje{
		Type:     "TEXT",
		Username: palabra[1],
		Text:     strings.Join(palabra[2:], " "),
	  }) 
     case "NEW_ROOM":
     	  if len(palabra) == 1 { // lo desconectamos y enviamos el porque
	     return
	  }
	  codificado.Encode(comun.Mensaje{
		Type:     "NEW_ROOM",
		Roomname: palabra[1],
	  })
     	  
     case "INVITE":
     	  if len(palabra) < 3 {
	     return // lo mismo de desconectarlo
	  }
	  usuarios := palabra[2:]
	  codificado.Encode(comun.Mensaje{
		Type:	   "INVITE",
		Roomname:  palabra[1],
		Usernames: usuarios,
	  })

     case "JOIN_ROOM":
     	  if len(palabra) == 1 {
	     return
	  }
	  codificado.Encode(comun.Mensaje{
		Type:	  "JOIN_ROOM",
		Roomname: palabra[1],
	  })

     case "ROOM_USERS":
     	  if len(palabra) == 1 {
	     return
	  }
	  codificado.Encode(comun.Mensaje{
		Type:	  "ROOM_USERS",
		Roomname: palabra[1],
	  })

     case "ROOM_TEXT":
     	  if len(palabra) < 3 {
	     return
	  }
	  msj := strings.Join(palabra[2:], " ")
	  codificado.Encode(comun.Mensaje{
		Type:	  "ROOM_TEXT",
		Roomname: palabra[1],
		Text:	  msj,
	  })

     case "LEAVE_ROOM":
     	  if len(palabra) < 2 {
	     return
	  }
	  codificado.Encode(comun.Mensaje{
		Type:	  "LEFT_ROOM",
		Romename: palabra[1],
	  })

     case "DISCONNECT": 
     	  codificado.Encode(comun.Mensaje{
		Type:	  "DISCONNECT",
	  })
	  return
     
     default: // msj para todos
	codificado.Encode(comun.Mensaje{
		Type: "PUBLIC_TEXT",
		Text: entrada,
	})
     }
}

// Función que procesa un mensaje
func ProcesaMensaje(mensaje comun.Mensaje) {
     switch mensaje.Type {
     case "NEW_USER":
     	  AvisoNuevoUsuario(mensaje)
	  
     case "NEW_STATUS":
     	  AvisoNuevoStatus(mensaje)
	  
     case "USER_LIST":
     	  ListaUsuarios(mensaje)
	  
     case "TEXT_FROM":
     	  TextoPrivado(mensaje)
	  
     case "PUBLIC_TEXT_FROM":
     	  TextoPublico(mensaje)

     case "INVITATION":
     	  Invitacion(mensaje)
	  
     case "JOINED_ROOM":
     	  UnidoASala(mensaje)

     case "ROOM_USER_LIST":
     	  ListaUsuariosSala(mensaje)

     case "ROOM_TEXT_FROM":
     	  TextoSala(mensaje)

     case "LEFT_ROOM":
     	  ClienteDejaSala

     case "DISCONNECTED":
     	  AvisoDesconectado(mensaje)

     case "RESPONSE":
     	  Respuesta(mensaje)
	  
     default:
	  MensajeDesconocido(mensaje)
	  
     }
}
