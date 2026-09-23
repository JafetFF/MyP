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
	fmt.Printf("%s: %s\n", mensaje.Username, mensaje.Text)
	fmt.Print("> ")
}

func Respuesta(mensaje comun.Mensaje) {
	fmt.Printf("Respuesta: %s  %s\n", mensaje.Operation, mensaje.Result)
}

func MensajeDesconocido(mensaje comun.Mensaje) {
	fmt.Printf("Mensaje no reconocido: %s\n", mensaje.Type)
}