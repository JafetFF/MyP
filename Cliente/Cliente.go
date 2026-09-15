package main

import "net"

type Cliente struct {
     direccion string
}

// Función que regresa un nuevo cliente
func NuevoCliente(direccion string) *Cliente {
     return &Cliente{
     	    direccion: direccion
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

     mensaje := mensaje{
     	 Type: "IDENTIFY",
	 Username: "Juan",
     }

     codificado := json.NewEncoder(conn)
     codificado.Encode(mensaje)


}
