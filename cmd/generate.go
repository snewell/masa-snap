package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/snewell/masa-snap/internal/masa-snap"
)

var (
	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate yaml for a new package",
		RunE:  doGenerate,
	}
)

func doGenerate(*cobra.Command, []string) error {
	if description == "" {
		rawBytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("error reading description: %v", err)
		}
		description = string(rawBytes)
	}

	snap := masasnap.SnapInfo{
		Name:        name,
		Version:     version,
		Summary:     summary,
		Description: description,
		Layout:      map[string]masasnap.SnapBind{},
	}

	for _, d := range bindDirs {
		ss := strings.Split(d, ":")
		snap.Layout[ss[0]] = masasnap.SnapBind{Bind: ss[1]}
	}
	for _, d := range bindFiles {
		ss := strings.Split(d, ":")
		snap.Layout[ss[0]] = masasnap.SnapBind{BindFile: ss[1]}
	}

	d, err := yaml.Marshal(&snap)
	if err != nil {
		return fmt.Errorf("error writing yaml: %v", err)
	}
	fmt.Print(string(d))
	return nil
}

func init() {
	addCommonArguments(generateCmd)
	generateCmd.PersistentFlags().StringVarP(&version, "version", "v", "1.0.0", "Package version")
	generateCmd.MarkPersistentFlagRequired("name")
	generateCmd.MarkPersistentFlagRequired("summary")

	rootCmd.AddCommand(generateCmd)
}
