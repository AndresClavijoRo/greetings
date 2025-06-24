# Saludos en go

Este es un ejemplo de un programa en Go que imprime saludos.

## Instalación

Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/AndresClavijoRo/greetings
```

## uso
Aquì tienes un ejemplo de como utilizar el paquete:

```go
package main

import (
	"fmt"
	"log"

	"github.com/AndresClavijoRo/greetings"
)

func main() {
	log.SetPrefix("greetings: ")

	names := []string{"Andres", "Maria", "Predro", "Gonzalez"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
```
este programa imprime un saludo personalizado para cada nombre en la lista `names`.