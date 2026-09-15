package main

import (
       "encoding/json"
       "net"
       "testing"
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
     var d mensaje
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