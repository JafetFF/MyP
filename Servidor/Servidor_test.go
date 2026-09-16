package main

import (
       "encoding/json"
       "net"
       "testing"
       "MyP/Comun"
       "sync"
)

// Prueba Unitaria utilizando net.Pipe()
// Se prueba el caso de un JSON inválido
func TestAtiendeConexion_JSONInvalido(t *testing.T) {
     cliente, servidor := net.Pipe()

     defer cliente.Close()

     servidorPrueba := &Servidor{}

     go servidorPrueba.AtiendeConexion(servidor)

     _, err := cliente.Write([]byte(`{
				       "type":"IDENTIFY",
				       "username":"Jorge"
				       ]`))
     if err != nil {
     	t.Fatalf("No se pudo escribir en el cliente: %v", err)
     }

     decodificado := json.NewDecoder(cliente)
     var d comun.Mensaje
     err = decodificado.Decode(&d)
     
     if err != nil {
     	t.Fatalf("No se pudo decodificar la respuesta: %v", err)
     }
     
     if d.Type != "RESPONSE" ||
        d.Operation != "INVALID" ||
	d.Result != "INVALID" {
	t.Errorf("Respuesta incorrecta")
     }     
}

// Prueba Unitaria utilizando net.Pipe()
// Se prueba el caso de un JSON con IDENTIFY inválido
func TestAtiendeConexion_IDENTIFY_INVALID(t *testing.T) {
     cliente, servidor := net.Pipe()

     defer cliente.Close()

     servidorPrueba := &Servidor{}

     go servidorPrueba.AtiendeConexion(servidor)

     _, err := cliente.Write([]byte(`{
				       "type":"NOTHING",
				       "username":"Marco"
				       }`))
     if err != nil {
     	t.Fatalf("No se pudo escribir en el cliente: %v", err)
     }

     decodificado := json.NewDecoder(cliente)
     var d comun.Mensaje
     err = decodificado.Decode(&d)
     
     if err != nil {
     	t.Fatalf("No se pudo decodificar la respuesta: %v", err)
     }
     
     if d.Type != "RESPONSE" ||
        d.Operation != "INVALID" ||
	d.Result != "NOT_IDENTIFIED" {
	t.Errorf("Respuesta incorrecta")
     }     
}


// Prueba Unitaria para verificar que se manda el mensaje de SUCCESS
func TestAtiendeConexion_SUCCESS(t *testing.T) {
     cliente, servidor := net.Pipe()

     defer cliente.Close()

     servidorPrueba := &Servidor{}

     go servidorPrueba.AtiendeConexion(servidor)

     _, err := cliente.Write([]byte(`{
				       "type":"IDENTIFY",
				       "username":"Nombre extravagante"
				       }`))
     if err != nil {
     	t.Fatalf("No se pudo escribir en el cliente: %v", err)
     }

     decodificado := json.NewDecoder(cliente)
     var d comun.Mensaje
     err = decodificado.Decode(&d)
     
     if err != nil {
     	t.Fatalf("No se pudo decodificar la respuesta: %v", err)
     }
     
     if d.Type != "RESPONSE" ||
        d.Operation != "IDENTIFY" ||
	d.Result != "SUCCESS" ||
	d.Extra != "Nombre extravagante" {
	t.Errorf("Respuesta incorrecta")
     }

     _, existe := servidorPrueba.clientes[d.Extra]
     if !existe {
     	t.Errorf("No se añadió el cliente")
     }
}

// Función para probar que no se permiten nombres repetidos
func TestAtiendeConexion_USER_ALREADY_EXISTS(t *testing.T) {
     cliente, servidor := net.Pipe()

     defer cliente.Close()

     servidorPrueba := &Servidor{
     		    clientes: make(map[string]ClienteConectado),
		    }


     servidorPrueba.clientes["Nombre repetido"] = ClienteConectado{
     	nombre:   "Nombre repetido",
	estado:   "ACTIVE",
	conexion: cliente,
     }

     go servidorPrueba.AtiendeConexion(servidor)

     _, err := cliente.Write([]byte(`{
				       "type":"IDENTIFY",
				       "username":"Nombre repetido"
				       }`))
     if err != nil {
     	t.Fatalf("No se pudo escribir en el cliente: %v", err)
     }

     decodificado := json.NewDecoder(cliente)
     var d comun.Mensaje
     err = decodificado.Decode(&d)
     
     if err != nil {
     	t.Fatalf("No se pudo decodificar la respuesta: %v", err)
     }
     
     if d.Type != "RESPONSE" ||
        d.Operation != "IDENTIFY" ||
	d.Result != "USER_ALREADY_EXISTS" ||
	d.Extra != "Nombre repetido" {
	t.Errorf("Respuesta incorrecta")
     }     
}


// Función para verificar que al desconectarse un cliente no está en la lista de clientes
func TestAtiendeConexion_DESCONEXION(t *testing.T) {
     cliente, servidor := net.Pipe()

     defer cliente.Close()

     servidorPrueba := &Servidor{}

     var espera sync.WaitGroup
     espera.Add(1)
     go func() {
     	defer espera.Done()
	servidorPrueba.AtiendeConexion(servidor)
     }()

     _, err := cliente.Write([]byte(`{
				       "type":"IDENTIFY",
				       "username":"José"
				       }`))
     if err != nil {
     	t.Fatalf("No se pudo escribir en el cliente: %v", err)
     }

     decodificado := json.NewDecoder(cliente)
     var d comun.Mensaje
     err = decodificado.Decode(&d)
     
     if err != nil {
     	t.Fatalf("No se pudo decodificar la respuesta: %v", err)
     }

     _, exist := servidorPrueba.clientes[d.Extra]
     if !exist {
     	t.Errorf("No se añadió el cliente")
     }

     

     cliente.Close()
     espera.Wait()

     _, existe := servidorPrueba.clientes[d.Extra]
     if existe {
     	t.Errorf("No se quitó el cliente al desconectar")
     }   
}

