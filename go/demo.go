package main

import (
	"fmt"
	"log"

	"github.com/popovicu/glua-kit/vm"
)

func main() {
	fmt.Println("Hello from Go")

	lua, err := vm.NewVm()

	if err != nil {
		log.Fatalf("unable to instantiate a Lua VM")
	}

	program := "print(\"Hello from Lua!\")"

	err = lua.RunCode(program)

	if err != nil {
		log.Fatalf("error running Lua: %w", err)
	}

	lua.Close()

	fmt.Println("Lua done")
}
