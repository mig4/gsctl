package endpoint

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/giantswarm/gscliauth/config"
	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"

	"github.com/giantswarm/gsctl/commands/errors"
	"github.com/giantswarm/gsctl/util"
)

const selectEndpointsCompletionFn = `
if [[ ${last_command} -eq "gsctl_select_endpoint" ]]; then
	__gsctl_get_endpoints;
fi
`

var (
	// Command performs the "select endpoint" function
	Command = &cobra.Command{
		Use:     "endpoint <endpoint>",
		Aliases: []string{"endpoints"},
		Short:   "Select endpoint to use",
		Long: `Select the API endpoint to use in subsequent commands.

To add an endpoint for the first time, or to re-login, use the 'gsctl login'
command with that endpoint.

To find out which endpoints are selectable, use the 'gsctl list endpoints'
command.
`,
		PreRun: selectEndpointPreRunOutput,
		Run:    selectEndpointRunOutput,
	}
)

func init() {
	util.SetCommandBashCompletion(&util.BashCompletionFunc{
		FnBody: selectEndpointsCompletionFn,
	})
}

// selectEndpointPreRunOutput does some pre-checks and, if necessary,
// shows output and exits.
func selectEndpointPreRunOutput(cmd *cobra.Command, cmdLineArgs []string) {
	err := verifySelectEndpointPreconditions(cmdLineArgs)

	if err == nil {
		return
	}

	var headline string
	var subtext string

	switch {
	case errors.IsEndpointMissingError(err):
		headline = "No endpoint specified."
		subtext = "Please give an endpoint URL to use. Use --help for details."
	default:
		headline = err.Error()
	}

	// print output
	fmt.Println(color.RedString(headline))
	if subtext != "" {
		fmt.Println(subtext)
	}
	os.Exit(1)
}

func verifySelectEndpointPreconditions(cmdLineArgs []string) error {
	if len(cmdLineArgs) == 0 {
		return microerror.Mask(errors.EndpointMissingError)
	}
	return nil
}

// selectEndpointRunOutput calls the actual function and
// shows the result of the command execution
func selectEndpointRunOutput(cmd *cobra.Command, cmdLineArgs []string) {
	err := config.Config.SelectEndpoint(cmdLineArgs[0])
	if err != nil {
		if config.IsEndpointNotDefinedError(err) {
			fmt.Println(color.RedString("The endpoint given is not defined."))
			fmt.Println("Please use 'gsctl login <email> -e <endpoint>' to add a new endpoint first.")
			os.Exit(1)
		}
		fmt.Println(color.RedString("Error: " + err.Error()))
	} else {
		fmt.Println(color.GreenString("Endpoint selected: %s", config.Config.SelectedEndpoint))
	}
}
