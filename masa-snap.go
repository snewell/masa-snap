package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"gopkg.in/yaml.v3"
)

type SnapBind struct {
	Bind     string `yaml:",omitempty"`
	BindFile string `yaml:"bind-file,omitempty"`
}

type SnapInfo struct {
	Name        string
	Version     string
	Summary     string
	Description string
	Layout      map[string]SnapBind `yaml:",omitempty"`
}

func main() {
	parser := argparse.NewParser("masa-snap", "Generate snapcraft yaml from user arguments")
	name := parser.String("n", "name", &argparse.Options{
		Required: true,
		Help:     "Name of the package",
	})
	version := parser.String("v", "version", &argparse.Options{
		Required: false,
		Help:     "Package version",
		Default:  "1.0.0",
	})
	summary := parser.String("s", "summary", &argparse.Options{
		Required: true,
		Help:     "Brief description of package",
	})
	description := parser.String("d", "description", &argparse.Options{
		Required: false,
		Help:     "Detailed description of package",
	})
	dirBinds := parser.StringList("b", "bind-dir", &argparse.Options{
		Required: false,
		Help:     "Bind a folder",
	})
	fileBinds := parser.StringList("f", "bind-file", &argparse.Options{
		Required: false,
		Help:     "Bind a file",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	if *description == "" {
		rawBytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("Error reading description: %v", err)
			os.Exit(1)
		}
		newDescription := string(rawBytes)
		description = &newDescription
	}

	snap := SnapInfo{
		Name:        *name,
		Version:     *version,
		Summary:     *summary,
		Description: *description,
		Layout:      map[string]SnapBind{},
	}

	for _, d := range *dirBinds {
		ss := strings.Split(d, ":")
		snap.Layout[ss[0]] = SnapBind{Bind: ss[1]}
	}
	for _, d := range *fileBinds {
		ss := strings.Split(d, ":")
		snap.Layout[ss[0]] = SnapBind{BindFile: ss[1]}
	}

	d, err := yaml.Marshal(&snap)
	if err != nil {
		fmt.Printf("Error writing yaml: %v", err)
	}
	fmt.Print(string(d))
}
