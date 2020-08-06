package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/julienbreux/clamp/functions"
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
		fmt.Fprintf(os.Stderr, "clamp: %s\n", err.Error())
	}
}

func run() error {
	in, err := input()
	if err != nil {
		return err
	}
	defer in.Close()

	var buf bytes.Buffer
	if err := transform(&buf, in, envVars()); err != nil {
		return fmt.Errorf("could not render: %w", err)
	}

	if _, err := io.Copy(os.Stdout, &buf); err != nil {
		return fmt.Errorf("could not print: %w", err)
	}

	return nil
}

func input() (*os.File, error) {
	switch flag.NArg() {
	case 0:
		return os.Stdin, nil
	case 1:
		filename := flag.Arg(0)
		f, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("unable to open %q: %w", filename, err)
		}
		return f, nil
	default:
		return nil, errors.New("incorrect usage: too many arguments")
	}
}

func transform(w io.Writer, r io.Reader, vars map[string]string) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("input is unreadable: %w", err)
	}

	tmpl, err := template.New("default").
		Option("missingkey=zero").
		Funcs(sprig.HermeticTxtFuncMap()).
		Funcs(functions.Map()).
		Parse(string(data))
	if err != nil {
		return err
	}

	return tmpl.Execute(w, vars)
}

func envVars() map[string]string {
	vars := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		vars[pair[0]] = pair[1]
	}
	return vars
}
