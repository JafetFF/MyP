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
     desconectado := make(chan struct{})
     go EscuchaServidor(conn, decodificado, desconectado)

     lector := bufio.NewReader(os.Stdin)
     entradas := make(chan string)

     go func() {
	for {
	    entrada, err := lector.ReadString('\n')
	    if err != nil {
	       return
	    }
	    entradas <- strings.TrimSpace(entrada)
	    }
     }()
     for {     	 	   
	fmt.Print("> ")
	select {
	case entrada := <-entradas:

	if entrada == "DISCONNECT" {
	   codificado.Encode(comun.Mensaje{
		Type: "DISCONNECT",
 	   })
	   return
	}
	
	InterpretaMensaje(entrada, codificado)

	case <-desconectado:
	     fmt.Println("\nEl servidor ha cerrado la conexión")
	     return
	}
     }
}

// Función que se mantiene a la espera de un mensaje nuevo
func EscuchaServidor(conn net.Conn, decodificado *json.Decoder, desconectado chan struct{}) {   
     for {
	 var m comun.Mensaje
     	 err := decodificado.Decode(&m)
	 if err != nil {
	    if errors.Is(err, net.ErrClosed) || err == io.EOF {
	       close(desconectado)
	       return
	    }
	    fmt.Printf("Error al recibir el mensaje %v\n", err)
	    close(desconectado)
	    return
	 }

	 ProcesaMensaje(m) 
     }
}

