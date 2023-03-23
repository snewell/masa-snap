package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	name        string
	version     string
	summary     string
	description string

	bindDirs  []string
	bindFiles []string

	rootCmd = &cobra.Command{
		Use:   "masa-snap",
		Short: "Manage yaml for snapcraft packages",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addCommonArguments(command *cobra.Command) {
	command.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the package")
	command.PersistentFlags().StringVarP(&summary, "summary", "s", "", "Brief description of package")
	command.PersistentFlags().StringVarP(&description, "description", "d", "", "Detailed description of package")

	command.PersistentFlags().StringSliceVarP(&bindDirs, "bind-dir", "b", []string{}, "Bind a folder")
	command.PersistentFlags().StringSliceVarP(&bindFiles, "bind-file", "f", []string{}, "Bind a file")
}

func AddCommand(command *cobra.Command) error {
	rootCmd.AddCommand(command)
	return nil
}
