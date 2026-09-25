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

func TextoPrivado(mensaje comun.Mensaje) {
	fmt.Printf("%s: %s\n", mensaje.Username, mensaje.Text)
}

func TextoPublico(mensaje comun.Mensaje) {
	fmt.Printf("%s: %s\n> ", mensaje.Username, mensaje.Text)
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
     default:

     }
}

func AvisoDesconectado(mensaje comun.Mensaje) {
     	fmt.Printf("Cliente desconectado: %s\n", mensaje.Username)
	fmt.Printf("> ")
}

func MensajeDesconocido(mensaje comun.Mensaje) {
	fmt.Printf("Mensaje no reconocido: %s\n", mensaje.Type)
}