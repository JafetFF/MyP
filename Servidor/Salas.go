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

// func InvitaCliente() {}

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

// func EnviaMensajeSala() {}

// func SalirSala() {}
