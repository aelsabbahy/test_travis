package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
func main() {
	cli.NewApp().Run(os.Args)
	fmt.Println(hello("world"))
}
