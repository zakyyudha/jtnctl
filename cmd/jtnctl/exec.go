package jtnctl

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zakyyudha/jtnctl/config"
	"github.com/zakyyudha/jtnctl/pkg/kubectl"
	"os"
)

var execCmd = &cobra.Command{
	Use:   "exec [pod-name] -- [command] [args...]",
	Short: "Execute a command in an active service",
	Long:  "Execute a command in an active Kubernetes service's container. This command is used to run arbitrary commands within a container.",
	Args:  cobra.MinimumNArgs(1), // Require at least one argument (pod-name)
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.LoadConfig()
		if err != nil {
			fmt.Printf("Error reading config file")
			os.Exit(1)
		}

		podName := args[0]
		namespace := config.Namespace
		command := args[1:]

		podToExec, err := kubectl.GetPodsName(namespace, podName)
		if err != nil {
			fmt.Printf("Error getting pod names: %v\n", err)
			os.Exit(1)
		}

		// Execute the command using kubectl exec
		kubectlArgs := []string{"-n", namespace, "exec", "-it", podToExec, "--"}
		kubectlArgs = append(kubectlArgs, command...)

		// Run 'kubectl logs' on the selected pod
		if err := kubectl.RunCommand(kubectlArgs...); err != nil {
			fmt.Printf("Error executing 'kubectl logs': %v\n", err)
			os.Exit(1)
		}
	},
}
