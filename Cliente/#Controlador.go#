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
     case "INVITE":
     case "JOIN_ROOM":
     case "ROOM_USERS":
     case "ROOM_TEXT":
     case "LEAVE_ROOM":
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
	  
     case "JOINED_ROOM":
     case "ROOM_USER_LIST":
     case "LEFT_ROOM":
     case "DISCONNECTED":
     	  AvisoDesconectado(mensaje)

     case "RESPONSE":
     	  Respuesta(mensaje)
	  
     default:
	  MensajeDesconocido(mensaje)
	  
     }
}
