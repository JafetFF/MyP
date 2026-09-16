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
     	log.Printf("Error al conectar: %v", err)
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
     log.Printf("mensaje: %s", m)
     
     go EscuchaServidor(conn, decodificado)

     lector := bufio.NewReader(os.Stdin)
     for{
	fmt.Print("> ")
	entrada, err := lector.ReadString('\n')
	if err != nil {
	   return
	}
	entrada = strings.TrimSpace(entrada)
	if entrada == "salir" {
	   codificado.Encode(comun.Mensaje{
		Type: "DISCONNECT",
 	   })
	   return
	}

	codificado.Encode(comun.Mensaje{
		Type: "PUBLIC_TEXT",
		Text: entrada,
	})
     }
}

// Función que se mantiene a la espera de un mensaje nuevo
func EscuchaServidor(conn net.Conn, decodificado *json.Decoder) {
     // decodificado = json.NewDecoder(conn)
     	 
     for {
	 var m comun.Mensaje
     	 err := decodificado.Decode(&m)
	 if err != nil {
	    log.Printf("Error al recibir el mensaje: %v", err)
	    return
	 }
	 log.Printf("Mensaje recibido: %+v", m)
	 ProcesaMensaje(m) 
     }
}

// Función que procesa un mensaje
func ProcesaMensaje(mensaje comun.Mensaje) {
     switch mensaje.Type {
     case "New_USER":
     case "NEW_STATUS":
     case "USER_LIST":
     case "TEXT_FROM":
     case "PUBLIC_TEXT_FROM":
     case "JOINED_ROOM":
     case "ROOM_USER_LIST":
     case "LEFT_ROOM":
     case "DISCONNECTED":
     case "RESPONSE":
     }

}
