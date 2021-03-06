package app

import (
	"github.com/jenkins-x/jx/cmd/codegen/generator"
	"github.com/jenkins-x/jx/cmd/codegen/util"
	"github.com/jenkins-x/jx/pkg/jx/cmd"
	"os"
	"path/filepath"

	"github.com/jenkins-x/jx/pkg/jx/cmd/templates"

	jxutil "github.com/jenkins-x/jx/pkg/util"

	"github.com/spf13/cobra"
)

// GenerateDocsOptions the options for the create client docs command
type GenerateDocsOptions struct {
	GenerateOptions
	ReferenceDocsVersion string
}

var (
	createClientDocsLong = templates.LongDesc(`This command code generates clients docs (Swagger,OpenAPI and HTML) for
	the specified custom resources.
 
`)

	createClientDocsExample = templates.Examples(`
		# lets generate client docs
		codegen docs
		
		# You will normally want to add a target to your Makefile that looks like:

		generate-clients-docs:
			codegen docs
		
		# and then call:

		make generate-clients-docs
`)
)

// NewCreateDocsCmd creates apidocs for CRDs
func NewCreateDocsCmd(commonOpts *cmd.CommonOptions) *cobra.Command {
	o := &GenerateDocsOptions{
		GenerateOptions: GenerateOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:     "docs",
		Short:   "Creates client docs for Custom Resources",
		Long:    createClientDocsLong,
		Example: createClientDocsExample,

		Run: func(c *cobra.Command, args []string) {
			o.Cmd = c
			o.Args = args
			err := run(o)
			cmd.CheckErr(err)
		},
	}

	wd, err := os.Getwd()
	if err != nil {
		util.AppLogger().Warnf("error getting working directory for %v\n", err)
	}

	cmd.Flags().StringVarP(&o.OutputBase, optionOutputBase, "o", filepath.Join(wd, "docs/apidocs"),
		"output base directory, by default the <current working directory>/docs/apidocs")
	return cmd
}

func run(o *GenerateDocsOptions) error {
	var err error
	if o.OutputBase == "" {
		return jxutil.MissingOption(optionOutputBase)
	}
	util.AppLogger().Infof("generating docs to %s\n", o.OutputBase)

	referenceDocsRepo, err := generator.InstallGenAPIDocs()
	if err != nil {
		return err
	}
	err = generator.GenerateAPIDocs(o.OutputBase)
	if err != nil {
		return err
	}
	err = generator.AssembleAPIDocsStatic(referenceDocsRepo, o.OutputBase)
	if err != nil {
		return err
	}
	err = generator.AssembleAPIDocs(o.OutputBase, filepath.Join(o.OutputBase, "site"))
	if err != nil {
		return err
	}
	return nil
}
