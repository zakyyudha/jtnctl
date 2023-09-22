package jtnctl

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zakyyudha/jtnctl/config"
	"github.com/zakyyudha/jtnctl/pkg/kubectl"
	"os"
)

var logsCmd = &cobra.Command{
	Use:   "logs [pod-name]",
	Short: "View logs of a Kubernetes pod",
	Long:  "View logs of a Kubernetes pod in the specified namespace. Use flags to control the output.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.LoadConfig()
		if err != nil {
			fmt.Printf("Error reading config file")
			os.Exit(1)
		}

		podName := args[0]
		namespace := config.Namespace

		podToFollow, err := kubectl.GetPodsName(namespace, podName)
		if err != nil {
			fmt.Printf("Error getting pod names: %v\n", err)
			os.Exit(1)
		}

		if len(podToFollow) == 0 {
			fmt.Printf("No pods found matching '%s' in namespace '%s'\n", podName, namespace)
			return
		}

		kubectlArgs := []string{"-n", namespace, "logs", podToFollow}

		// Get the flags passed to the jtnctl logs command
		flags := cmd.Flags()
		follow, _ := flags.GetBool("follow")
		tail, _ := flags.GetInt("tail")

		if follow {
			kubectlArgs = append(kubectlArgs, "-f")
		}
		if tail > 0 {
			kubectlArgs = append(kubectlArgs, fmt.Sprintf("--tail=%d", tail))
		}

		// Execute the command using kubectl logs
		fmt.Println(kubectlArgs)

		// Run 'kubectl logs' on the selected pod
		if err := kubectl.RunCommand(kubectlArgs...); err != nil {
			fmt.Printf("Error executing 'kubectl logs': %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	// Add flags for the "logs" command
	logsCmd.Flags().BoolP("follow", "f", false, "Specify if you want to follow the logs")
	logsCmd.Flags().Int("tail", 0, "Number of lines to show from the end of the logs")
}
