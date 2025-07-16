package app

import (
	"github.com/spf13/cobra"

	"github.com/go-cheetah/cheetah/cmd/cheetah/app/options"
)

var Version = ""

func NewServerCommand() *cobra.Command {
	o := options.NewAppOptions()
	cmd := &cobra.Command{
		Use:  "cheetah",
		Long: `Long describe.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		DisableAutoGenTag: true,
	}
	cmd.CompletionOptions.HiddenDefaultCmd = true

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of cheetah.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(Version)
		},
	}
	projectCmd := &cobra.Command{
		Use:   "create",
		Short: "Create type project.",
		Long:  "Create type project. type is 'ansible'、'gitbook'、'mdbook'、'command'、'http'、'mvc'、'grpc'、'simple'.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(projectCmd)

	projectCmd.PersistentFlags().StringVarP(&o.Name, "name", "n", "demo", "Project name.")
	projectCmd.PersistentFlags().StringVarP(&o.Path, "path", "p", "./", "Project path.")

	newCommand(o, projectCmd, "ansible")
	newCommand(o, projectCmd, "mvc")
	newCommand(o, projectCmd, "gitbook")
	newCommand(o, projectCmd, "mdbook")
	newCommand(o, projectCmd, "simple")
	newCommand(o, projectCmd, "http")
	newCommand(o, projectCmd, "command")
	newCommand(o, projectCmd, "grpc")

	return cmd
}

func newCommand(o *options.AppOptions, projectCmd *cobra.Command, typeName string) {
	tmpCmd := &cobra.Command{
		Use:   typeName,
		Short: "Create " + typeName + " project.",
		RunE: func(cmd *cobra.Command, args []string) error {
			o.Type = typeName
			return run(o)
		},
	}
	projectCmd.AddCommand(tmpCmd)
}

func run(o *options.AppOptions) (err error) {
	server, err := o.NewServer()
	server.Run()
	return
}
