package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Para trabajar con json el flujo es myVarJSON, err := json.Marshall(myVar) y handlear el error, despues si se quiere usar como string es string(myVarJSON)

func main() {

	var (
		john = Person{
			ID:          123,
			Name:        "John Doe",
			DateOfBirth: time.Date(1998, 02, 21, 7, 30, 20, 100, time.Local),
		}

		johnEmp = Employee{
			ID:       456,
			Position: "Manager",
			Person:   john,
		}
	)

	johnJSON, errJohn := json.Marshal(john)
	johnEmpJSON, errJohnEmp := json.Marshal(johnEmp)

	fmt.Println("John Doe es una instancia de Persona: ")
	if errJohn != nil {
		fmt.Println(errJohn)
	} else {
		fmt.Println(string(johnJSON))
	}

	fmt.Println("John Doe pego laburo y ahora es una instancia de Empleado tambien: ")
	if errJohnEmp != nil {
		fmt.Println(errJohnEmp)
	} else {
		fmt.Println(string(johnEmpJSON))
	}
}
