package main

import (
       "os"
       "net"
)
func main() {

     var direccion string
     var nombre    string
     for _, arg := range os.Args[1:] {
     	  if _, _, err := net.SplitHostPort(arg); err == nil {
	     direccion = arg
	  } else {
	     nombre = arg
	  }
     }
     cliente := NuevoCliente(direccion, nombre)
     cliente.Conecta()
}
