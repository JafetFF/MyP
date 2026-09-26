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
     	  mensaje := comun.Mensaje{
		Type:   "STATUS",
		Status: "",
	  }
     	  if len(palabra) != 1 {
	     mensaje.Status = palabra[1]
	  }
	  codificado.Encode(mensaje)
	  
     case "TEXT": // msj privado
     	  mensaje := comun.Mensaje{
		Type:     "TEXT",
		Username: "",
		Text:     "",
	  }
     	  if len(palabra) > 2 {
	     mensaje.Username = palabra[1]
	     mensaje.Text     = strings.Join(palabra[2:], " ")
	  }
	  codificado.Encode(mensaje)
	  
     case "NEW_ROOM":
     	  mensaje := comun.Mensaje{
		Type:     "NEW_ROOM",
		Roomname: "",
	  }
     	  if len(palabra) != 1 {
	     mensaje.Roomname = palabra[1]
	  }
	  codificado.Encode(mensaje)
     	  
     case "INVITE":
     	  mensaje := comun.Mensaje{
		Type:      "INVITE",
		Roomname:  "",
		Usernames: []string{},
	  }
     	  if len(palabra) > 2 {
	     mensaje.Roomname  = palabra[1]
	     mensaje.Usernames = palabra[2:]
	  }
	  codificado.Encode(mensaje)

     case "JOIN_ROOM":
     	  mensaje := comun.Mensaje{
		Type:     "JOIN_ROOM",
		Roomname: "",
	  }
     	  if len(palabra) != 1 {
	     mensaje.Roomname = palabra[1]
	  }
	  codificado.Encode(mensaje)

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
		Type:	  "LEAVE_ROOM",
		Roomname: palabra[1],
	  })
     
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
     	  ClienteDejaSala(mensaje)

     case "DISCONNECTED":
     	  AvisoDesconectado(mensaje)

     case "RESPONSE":
     	  Respuesta(mensaje)
	  
     default:
	  MensajeDesconocido(mensaje)
	  
     }
}
