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

     if m.Result != "SUCCESS" {
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
     	 
     for {
	 var m comun.Mensaje
     	 err := decodificado.Decode(&m)
	 if err != nil {
	    log.Printf("Error al recibir el mensaje: %v", err)
	    return
	 }

	 ProcesaMensaje(m) 
     }
}

// Función que procesa un mensaje
func ProcesaMensaje(mensaje comun.Mensaje) {
     switch mensaje.Type {
     case "NEW_USER":
     	  fmt.Printf("Nuevo usuario conectado: %s\n", mensaje.Username)
	  
     case "NEW_STATUS":
     	  fmt.Printf("%s ha cambiado a %s\n", mensaje.Username, mensaje.Status)
	  
     case "USER_LIST":
     	  fmt.Printf("Lista de usuarios:")
	  for usuario, estado := range mensaje.Users{
	      	       fmt.Printf("%s: %s\n", usuario, estado)
	  }
	  
     case "TEXT_FROM":
     	  fmt.Printf("%s: %s\n", mensaje.Username, mensaje.Text)
	  
     case "PUBLIC_TEXT_FROM":
     	  fmt.Printf("%s: %s\n", mensaje.Username, mensaje.Text)
	  fmt.Print("> ")
	  
     case "JOINED_ROOM":
     case "ROOM_USER_LIST":
     case "LEFT_ROOM":
     case "DISCONNECTED":
     case "RESPONSE":
     	  fmt.Printf("Respuesta: %s - %s\n", mensaje.Operation, mensaje.Result)
     default:
	fmt.Printf("Mensaje no reconocido: %s\n", mensaje.Type)
	  
     }

}
