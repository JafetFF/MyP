package main

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

func (s *Sala) AgregaCliente(nombre) {

}

// func GetUsuarios() {}

// func EnviaMensajeSala() {}

// func SalirSala() {}
