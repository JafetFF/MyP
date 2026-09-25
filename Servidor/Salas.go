package main

import (
       "MyP/Comun"
       "encoding/json"
       "fmt"
       "net"
)

type Sala struct {
     nombre     string
     listaUsers map[string]*ClienteConectado
     invitados  map[string]bool
}

func NuevaSala(nombre string) *Sala {
     return &Sala{
     	    nombre:	nombre,
	    listaUsers: make(map[string]*ClienteConectado),
	    invitados:  make(map[string]bool),
     }
}

func (s *Servidor) InvitaClientes(m comun.Mensaje, nombre string, usuarios []string) {
     for _, usuario := range usuarios {
	 respuesta := comun.Mensaje{
	      Type:     "INVITATION",
	      Username: nombre,
	      Roomname: m.Roomname,
	 }
	 data, err := json.Marshal(respuesta)
	 if err != nil {
	    fmt.Printf("Error al recibir mensaje: %v\n", err)
	 }
	 
	 fmt.Printf(">>> %s\n", data)
	 cliente, _ := s.clientes[usuario]
	 codificador := json.NewEncoder(cliente.conexion)
	 codificador.Encode(respuesta)
     }
}

func (s *Sala) AgregaCliente(cliente *ClienteConectado) {
     s.listaUsers[cliente.nombre] = cliente
}

// Función que regresa la lista de clientes de una sala
func (s *Sala) ListaSala(conn net.Conn) {
     users := make(map[string]string)
     for nombre, cliente := range s.listaUsers {
	      users[nombre] = cliente.estado
     }
     respuesta := comun.Mensaje{
 	    Type:     "ROOM_USER_LIST",
	    Roomname: s.nombre,
	    Users:    users,
     }
     data, err := json.Marshal(respuesta)
     if err != nil {
          fmt.Printf("Error al recibir mensaje: %v\n", err)
     }
	 
     fmt.Printf(">>> %s\n", data)
     codificado := json.NewEncoder(conn)
     codificado.Encode(respuesta)
}

func (s *Sala) ContieneCliente(nombre string) bool {
     _, existe := s.listaUsers[nombre]
     return existe
}

// Une a un cliente a una sala
func (s *Servidor) UneCliente(m comun.Mensaje, nombre string, sala *Sala) {
     respuesta := comun.Mensaje{
     	       Type:      "RESPONSE",
	       Operation: "JOIN_ROOM",
	       Result:    "SUCCESS",
	       Extra:  	  m.Roomname,
     }
     data, err := json.Marshal(respuesta)
     if err != nil {
     	fmt.Printf("Error al recibir mensaje: %v\n", err)
     }
	 
     fmt.Printf(">>> %s\n", data)

     cliente, _ := s.clientes[nombre]
     sala.listaUsers[nombre] = cliente
     codificador := json.NewEncoder(cliente.conexion)
     codificador.Encode(respuesta)
}

// func EnviaMensajeSala() {}

// func SalirSala() {}
