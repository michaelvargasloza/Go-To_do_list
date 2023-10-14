package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tasks := []string{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Lista de tareas:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}

		fmt.Println("Opciones:")
		fmt.Println("1. Agregar tarea")
		fmt.Println("2. Completar tarea")
		fmt.Println("3. Salir")

		fmt.Print("Seleccione una opción: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Print("Ingrese la tarea que desea agregar: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			tasks = append(tasks, task)
		case "2":
			fmt.Print("Ingrese el número de la tarea que desea marcar como completada: ")
			var taskNum int
			_, err := fmt.Scanf("%d", &taskNum)
			if err != nil || taskNum < 1 || taskNum > len(tasks) {
				fmt.Println("Número de tarea no válido.")
				continue
			}
			tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)
		case "3":
			fmt.Println("Saliendo del programa.")
			os.Exit(0)
		default:
			fmt.Println("Opción no válida. Inténtelo de nuevo.")
		}
	}
}
