package main

import (
	"mysql-driver-go/handlers"
	// "mysql-driver-go/models"
)

func main() {

	// cliente := models.Cliente{Nombre: "Nicola", Correo: "test@testl.com", Telefono: "00000"}
	// handlers.Editar(2, cliente)
	// handlers.Insertar(cliente)
	// handlers.ListarPorId(2)
	//    handlers.Eliminar(2)
	// handlers.Listar()
	handlers.Ejecutar()
}
