package main

import (
       "MyP/Comun"
       "fmt"
)

func AvisoNuevoUsuario(mensaje comun.Mensaje) {
     fmt.Printf("Nuevo usuario conectado: %s\n", mensaje.Username)
	  fmt.Print("> ")
}

func AvisoNuevoStatus(mensaje comun.Mensaje) {
     fmt.Printf("%s ha cambiado su estado a %s\n", mensaje.Username, mensaje.Status)
     fmt.Print("> ")
}

func ListaUsuarios(mensaje comun.Mensaje) {
     fmt.Printf("Lista de usuarios:\n")
     for usuario, estado := range mensaje.Users{
     	 fmt.Printf("> %s: %s\n", usuario, estado)
     }
     fmt.Print("> ")
}

func ListaUsuariosSala(mensaje comun.Mensaje) {
     fmt.Printf("Lista de usuarios de la sala %s:\n", mensaje.Roomname)
     for usuario, estado := range mensaje.Users{
     	 fmt.Printf("> %s: %s\n", usuario, estado)
     }
     fmt.Print("> ")
}

func Invitacion(mensaje comun.Mensaje) {
     fmt.Printf("\n	Te ha invitado %s a unirte a la sala %s\n> ", mensaje.Username, mensaje.Roomname)
}

func TextoPrivado(mensaje comun.Mensaje) {
     fmt.Printf("%s: %s\n", mensaje.Username, mensaje.Text)
}

func TextoPublico(mensaje comun.Mensaje) {
     fmt.Printf("%s: %s\n> ", mensaje.Username, mensaje.Text)
}

func TextoSala(mensaje comun.Mensaje) {
     fmt.Printf("<%s> %s: %s\n> ", mensaje.Roomname, mensaje.Username, mensaje.Text)
}

func Respuesta(mensaje comun.Mensaje) {
     switch mensaje.Operation {
     case "TEXT":
     	  fmt.Printf("RESPUESTA: No se envió el mensaje porque el cliente %s no existe\n> ", mensaje.Extra)
     case "IDENTIFY":
     	  fmt.Printf("RESPUESTA: Elije otro nombre. Ya existe un cliente llamado %s\n", mensaje.Extra)
     case "INVALID":
     	  if mensaje.Result == "NOT_IDENTIFIED" {
	     fmt.Printf("RESPUESTA: Debes identificarte con un nombre\n")
	  }
	  fmt.Println("RESPUESTA: el JSON es inválido o no se recibieron los valores esperados")
     case "NEW_ROOM":
     	  fmt.Printf("\n	Se creó correctamente la sala %s\n> ", mensaje.Extra)
     case "ROOM_USERS":
     	  if mensaje.Result == "NO_SUCH_ROOM" {
     	     fmt.Printf("RESPUESTA: La sala %s no existe\n> ", mensaje.Extra)
	  }
	  if mensaje.Result == "NOT_JOINED" {
	     fmt.Printf("RESPUESTA: No puedes pedir la lista ni invitar a otros clientes a la sala %s porque no te has unido a ella\n> ", mensaje.Extra)
	  }
     case "INVITE":
       	  fmt.Printf("RESPUESTA: No puedes invitar a %s a unirse a una sala, pues %s no existe\n> ", mensaje.Extra, mensaje.Extra)
     case "JOIN_ROOM":
     	  if mensaje.Result == "SUCCESS" {
     	     fmt.Printf("	Te has unido a la sala %s\n> ", mensaje.Roomname)
	  }
	  if mensaje.Result == "NO_SUCH_ROOM" {
	     fmt.Printf("RESPUESTA: La sala %s no existe\n> ", mensaje.Extra)
	  }
	  if mensaje.Result == "NOT_INVITED" {
	     fmt.Printf("RESPUESTA: No te puedes unir a la sala %s porque no has recibido una invitación\n> ", mensaje.Extra)
	  }

     case "ROOM_TEXT":
     	  if mensaje.Result == "NO_SUCH_ROOM" {
	     fmt.Printf("RESPUESTA: La sala %s no existe\n> ", mensaje.Extra)
	  }
	  if mensaje.Result == "NOT_JOINED" {
	     fmt.Printf("RESPUESTA: No puedes enviar un mensaje a la sala %s porque no te has unido a ella\n> ", mensaje.Extra)
	  }
     case "LEAVE_ROOM":
     	  if mensaje.Result == "NO_SUCH_ROOM" {
	     fmt.Printf("RESPUESTA: La sala %s no existe\n> ", mensaje.Extra)
	  }
	  if mensaje.Result == "NOT_JOINED" {
	     fmt.Printf("RESPUESTA: No puedes salir de la sala %s porque no te has unido a ella\n> ", mensaje.Extra)
	  }
     default:

     }
}

func AvisoDesconectado(mensaje comun.Mensaje) {
     fmt.Printf("Cliente desconectado: %s\n", mensaje.Username)
     fmt.Printf("> ")
}

func UnidoASala(mensaje comun.Mensaje) {
     fmt.Printf("Se ha unido %s a la sala %s\n> ", mensaje.Username, mensaje.Roomname)
}

func MensajeDesconocido(mensaje comun.Mensaje) {
     fmt.Printf("Mensaje no reconocido: %s\n", mensaje.Type)
}

func ClienteDejaSala(m comun.Mensaje) {
     fmt.Printf("%s salió de la sala %s\n> ", m.Username, m.Roomname)
}