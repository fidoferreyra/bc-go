package main

import (
	"errors"
	"fmt"
)

const (
	Nombre       = "NOMBRE"
	Apellido     = "APELLIDO"
	Dni          = "DNI"
	FechaIngreso = "FECHA_INGRESO"
)

type Estudiante struct {
	Id           int    `json:"-"`
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	Dni          string `json:"dni"`
	FechaIngreso string `json:"fecha_ingreso"`
}

func NuevoEstudiante(id int, nombre, apellido, dni, fechaIngreso string) *Estudiante {
	return &Estudiante{
		Id:           id,
		Nombre:       nombre,
		Apellido:     apellido,
		Dni:          dni,
		FechaIngreso: fechaIngreso,
	}
}

func Add(e *Estudiante) {
	Estudiantes = append(Estudiantes, e)
	fmt.Println("Se agrego al estudiante:")
	e.Detalle()
}

func (e Estudiante) Detalle() {
	fmt.Printf("Nombre : %v \n", e.Nombre)
	fmt.Printf("Apellido : %v \n", e.Apellido)
	fmt.Printf("Dni : %v \n", e.Dni)
	fmt.Printf("Fecha de Ingreso : %v \n", e.FechaIngreso)
}

func (e *Estudiante) editarAtributo(atributo, valor string) error {
	switch atributo {
	case Nombre:
		e.Nombre = valor
		return nil
	case Apellido:
		e.Apellido = valor
		return nil
	case Dni:
		e.Dni = valor
		return nil
	case FechaIngreso:
		e.FechaIngreso = valor
		return nil
	default:
		return errors.New("el atributo seleccionado no existe")
	}
}

func Editar(id int, atributo, valor string) error {
	for _, estudiante := range Estudiantes {
		if estudiante.Id == id {
			return estudiante.editarAtributo(atributo, valor)
		}
	}
	return errors.New("el id ingresado no corresponde a un estudiante")
}
