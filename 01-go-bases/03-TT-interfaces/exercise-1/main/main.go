package main

import "fmt"

var Estudiantes = []*Estudiante{}

func main() {
	var (
		juan    = NuevoEstudiante(1, "Juan", "Pérez", "12345678", "2020-01-01")
		maria   = NuevoEstudiante(2, "María", "Gómez", "23456789", "2021-02-01")
		luis    = NuevoEstudiante(3, "Luis", "Rodríguez", "34567890", "2019-03-01")
		ana     = NuevoEstudiante(4, "Ana", "Fernández", "45678901", "2022-04-01")
		roberto = NuevoEstudiante(5, "Roberto", "Sánchez", "56789012", "2018-05-01")
	)

	Add(juan)
	Add(luis)
	Add(maria)
	Add(ana)
	Add(roberto)

	Editar(juan.Id, Nombre, "Carlos")
	fmt.Printf("Ahora el estudiante Juan se llama %v \n", juan.Nombre)
}
