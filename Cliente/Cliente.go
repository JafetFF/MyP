package main

import (
       "net"
       "encoding/json"
       "log"
       "MyP/Comun"
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
     
     log.Printf("tipo: %s\n", m.Type)
     log.Printf("operacion: %s\n", m.Operation)
     log.Printf("resultado: %s\n", m.Result)
     log.Printf("extra: %s\n", m.Extra)
     
     
}
