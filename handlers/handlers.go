package handlers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	// "log"
	"mysql-driver-go/conectar"
	"mysql-driver-go/models"
)

func Listar() {

	conectar.ConnectToDB()

	sql := "select id, nombre, correo, telefono from clientes order by id desc;"

	datos, err := conectar.MySqlDatabase.Query(sql)

	if err != nil {
		fmt.Println(err)
	}

	defer conectar.MySqlDatabase.Close()

	clientes := models.Clientes{}

	for datos.Next() {
		dato := models.Cliente{}
		datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		clientes = append(clientes, dato)
	}

	fmt.Println(clientes)
}

func ListarPorId(id int) {
	conectar.ConnectToDB()

	sql := "select id, nombre, correo, telefono from clientes where id=?;"

	datos, err := conectar.MySqlDatabase.Query(sql, id)

	if err != nil {
		fmt.Println(err)
	}

	defer conectar.MySqlDatabase.Close()

	// var dato models.Cliente
	// // aunque venga un solo dto por id sera una lista lo que devuelve la peticion sql
	// // por loq ue la recorro con for y añado a la variable dato de cliente con el scan
	// // arruba lo hice con append cuando es un arreglo lo que tendre
	// for datos.Next() {
	// 	datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
	// }
	// fmt.Println(dato)

	for datos.Next() {
		var dato models.Cliente
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Id: %v | Nombre: %v | Telefono: %v | Email: %v \n",
			dato.Id, dato.Nombre, dato.Telefono, dato.Correo)
	}

}

func Insertar(cliente models.Cliente) {

	conectar.ConnectToDB()

	sql := "insert into clientes values(null, ?,?,?);"
	result, err := conectar.MySqlDatabase.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println("Se creo el registro exitosamente")

	defer conectar.MySqlDatabase.Close()
}

func Editar(id int, cliente models.Cliente) {

	conectar.ConnectToDB()
	sql := "update clientes set nombre=?, correo=?, telefono=? where id=?;"
	result, err := conectar.MySqlDatabase.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, id)

	if err != nil {
		panic(err)
	}

	fmt.Println("Se edito el registro exitosamente", result)
}

func Eliminar(id int) {
	conectar.ConnectToDB()

	sql := "delete from clientes where id=?"
	_, err := conectar.MySqlDatabase.Exec(sql, id)

	if err != nil {
		panic("Error al eliminar un cliente")
	}
	fmt.Println("se elimino el resgistro exitosamente")
}

// funciones de trabajo
var ID int
var nombre, correo, telefono string

func Ejecutar() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Seleccione una opcion : \n\n\n 1. Listar clientes \n\n 2. Listar cliente por ID \n\n 3. Crear cliente \n\n 4. Editar cliente \n\n 5. Eliminar cliente \n\n")

	if scanner.Scan() {
		for {
			if scanner.Text() == "1" {
				Listar()
				return
			}
			if scanner.Text() == "2" {
				fmt.Println("Ingrese el ID del cliente : \n")
				var idCliente int
				if scanner.Scan() {

					idCliente, _ = strconv.Atoi(scanner.Text())

				}
				ListarPorId(idCliente)
				return
			}
			if scanner.Text() == "3" {
				fmt.Println("Ingrese nombre : \n")
				if scanner.Scan() {
					nombre = scanner.Text()

				}
				fmt.Println("Ingrese correo : \n")

				if scanner.Scan() {
					correo = scanner.Text()

				}
				fmt.Println("Ingrese telefono : \n")

				if scanner.Scan() {
					telefono = scanner.Text()

				}

				cliente := models.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
				Insertar(cliente)
				return
			}
			if scanner.Text() == "4" {

				fmt.Println("Ingrese el ID del cliente : \n")
				var idCliente int
				if scanner.Scan() {

					idCliente, _ = strconv.Atoi(scanner.Text())

				}

				fmt.Println("Ingrese nombre : \n")
				if scanner.Scan() {
					nombre = scanner.Text()

				}
				fmt.Println("Ingrese correo : \n")

				if scanner.Scan() {
					correo = scanner.Text()

				}
				fmt.Println("Ingrese telefono : \n")

				if scanner.Scan() {
					telefono = scanner.Text()

				}
				cliente := models.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
				Editar(idCliente, cliente)
				return

			}

			if scanner.Text() == "5" {
				fmt.Println("Ingrese el ID del cliente : \n")
				var idCliente int
				if scanner.Scan() {

					idCliente, _ = strconv.Atoi(scanner.Text())

				}

				Eliminar(idCliente)
				return
			}

		}
	}
}
