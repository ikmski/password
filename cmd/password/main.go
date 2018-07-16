package main

import (
	"fmt"

	"github.com/ikmski/password"
	"github.com/urfave/cli"
)

var (
	version  string
	revision string
)

func verifyPassword(c *cli.Context) error {

	if c.NArg() != 1 {
		return fmt.Errorf("A password is required as an argument")
	}
	pass := c.Args()[0]

	pp := password.Default()

	ok, _ := pp.Verify(pass)
	if !ok {
		return fmt.Errorf("The password is not Valid")
	}

	return nil
}

func generatePassword(c *cli.Context) error {

	pp := password.Default()

	pass := pp.Random()
	fmt.Println(pass)

	return nil
}

func main() {

	app := cli.NewApp()
	app.Name = "password"
	app.Usage = "Password"
	app.Description = "Validate the password based on the policy. And generate a password"
	app.Version = version

	app.Commands = []cli.Command{
		{
			Name:    "verify",
			Aliases: []string{"v"},
			Usage:   "verify password",
			Action:  verifyPassword,
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate password",
			Action:  generatePassword,
		},
	}

	app.RunAndExitOnError()
}
