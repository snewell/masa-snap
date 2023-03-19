package main

import (
	"fmt"
	"io"
	"os"

	"github.com/akamensky/argparse"
	"gopkg.in/yaml.v3"
)

type SnapInfo struct {
	Name        string
	Version     string
	Summary     string
	Description string
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
	}
	d, err := yaml.Marshal(&snap)
	if err != nil {
		fmt.Printf("Error writing yaml: %v", err)
	}
	fmt.Print(string(d))
}
