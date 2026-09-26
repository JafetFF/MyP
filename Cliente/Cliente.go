package main

import (
       "net"
       "fmt"
       "encoding/json"
       "log"
       "MyP/Comun"
       "bufio"
       "os"
       "strings"
       "io"
       "errors"
)

type Cliente struct {
     direccion string
     nombre    string
}

// Función que regresa un nuevo cliente
func NuevoCliente(direccion string, nombre string) *Cliente {
     return &Cliente{
     	    direccion: direccion,
	    nombre:    nombre,
     }
}

// Función que conecta el cliente con el servidor
func (c *Cliente) Conecta() {
  
     conn, err := net.Dial("tcp", c.direccion)

     if err != nil {
     	fmt.Printf("Debes poner una IP:HOST válida\n")
	return
     }
     defer conn.Close()

     mensaje := comun.Mensaje{
     	 Type: "IDENTIFY",
	 Username: c.nombre,
     }

     codificado := json.NewEncoder(conn)
     codificado.Encode(mensaje)

     decodificado := json.NewDecoder(conn)
     var m comun.Mensaje
     err = decodificado.Decode(&m)

     if err != nil {
     	log.Printf("Error al recibir respuestaa: %v", err)
	return
     }

     if m.Result != "SUCCESS" {
     	ProcesaMensaje(m)
     	return
     }
     
     go EscuchaServidor(conn, decodificado)

     lector := bufio.NewReader(os.Stdin)
     for{
	     
	fmt.Print("> ")
	entrada, err := lector.ReadString('\n')
	if err != nil {
	   return
	}
	entrada = strings.TrimSpace(entrada)
	if entrada == "DISCONNECT" {
	   codificado.Encode(comun.Mensaje{
		Type: "DISCONNECT",
 	   })
	   return
	}
	
	InterpretaMensaje(entrada, codificado)

     }
}

// Función que se mantiene a la espera de un mensaje nuevo
func EscuchaServidor(conn net.Conn, decodificado *json.Decoder) {   
     for {
	 var m comun.Mensaje
     	 err := decodificado.Decode(&m)
	 if err != nil {
	    if errors.Is(err, net.ErrClosed) || err == io.EOF {
	       return
	    }
	    fmt.Printf("Error al recibir el mensaje %v\n", err)
	    return
	 }

	 ProcesaMensaje(m) 
     }
}

