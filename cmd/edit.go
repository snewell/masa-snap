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
	editCmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit yaml for an existing package",
		RunE:  doEdit,
	}
)

type editPair struct {
	first  *string
	second *string
}

func doEdit(*cobra.Command, []string) error {
	snap := masasnap.SnapInfo{}

	rawBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(rawBytes, &snap)
	if err != nil {
		return err
	}

	edits := []editPair{
		editPair{first: &name, second: &snap.Name},
		editPair{first: &version, second: &snap.Version},
		editPair{first: &summary, second: &snap.Summary},
		editPair{first: &description, second: &snap.Description},
	}

	for i := range edits {
		if *edits[i].first != "" {
			*edits[i].second = *edits[i].first
		}
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
	addCommonArguments(editCmd)
	editCmd.PersistentFlags().StringVarP(&version, "version", "v", "", "Package version")

	rootCmd.AddCommand(editCmd)
}
