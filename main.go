package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

var WastebandVersion = "0.1"
var conf *configuration

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	log.SetOutput(ioutil.Discard)

	conf, err := getConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(*conf)

	c := &cli.CLI{
		Args:     os.Args[:1],
		Name:     "wasteband",
		Version:  WastebandVersion,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("wasteband"),
	}

	_, err = c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return err
	}

	return nil
}
