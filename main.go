package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/julienbreux/clamp/version"
)

var (
	showVersion = false
)

func init() {
	flag.BoolVar(&showVersion, "v", false, "Print version information and quit")
}

func main() {
	flag.Parse()

	if showVersion {
		version.Print()
		return
	}

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}
}

func run() error {
	input()
	return nil
}

func input() {
	var data []byte
	var err error

	switch flag.NArg() {
	case 0:
		file := os.Stdin
		fi, err := file.Stat()
		check(err)
		if fi.Size() == 0 {
			printInputErrorMessage()
		}

		data, err = ioutil.ReadAll(os.Stdin)
		check(err)
		err = transform(data)
		check(err)
		break
	case 1:
		data, err = ioutil.ReadFile(flag.Arg(0))
		check(err)
		err = transform(data)
		check(err)
		break
	default:
		printInputErrorMessage()
	}
}

func transform(data []byte) error {
	tmpl, err := template.New("default").
		Option("missingkey=zero").
		Funcs(sprig.HermeticTxtFuncMap()).
		Parse(string(data))

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, envVars())
}

func envVars() map[string]string {
	vars := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		vars[pair[0]] = pair[1]
	}
	return vars
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printInputErrorMessage() {
	fmt.Printf("Input must be from stdin or file and non empty\n")

	os.Exit(1)
}
