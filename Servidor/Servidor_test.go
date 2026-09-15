package main

import (
       "encoding/json"
       "net"
       "testing"
       "MyP/Comun"
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