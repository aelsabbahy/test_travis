package main

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

func hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
func main() {
	cli.NewApp().Run(os.Args)
	fmt.Println(hello("world"))
	h, err := homedir.Dir()
	fmt.Println("homedir:", h, err)
	if h, err := homedir.Dir(); err != nil {
		fmt.Println("homedir:", h, err)
	} else {
		cwd, err := os.Getwd()
		if err == nil {
			homeDir := filepath.Join(cwd, ".glide")
			fmt.Println("cwd:", homeDir, err)
		} else {
			homeDir := ".glide"
			fmt.Println("else:", homeDir, err)
		}
	}
}
