package main

import "fmt"

func main() {
     var nombre string
     fmt.Print("Escribe tu nombre: ")
     fmt.Scanln(&nombre)
     cliente := NuevoCliente("localhost:8080", nombre)
     cliente.Conecta()
}
