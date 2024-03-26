package main

import (
	"fmt"
	"sync"
)

type Config struct {
	PORT int
	ENV  string
}

type Application struct {
	config Config
	wg     sync.WaitGroup
}

func main() {
	fmt.Println("Hola amigos!!!")
}
