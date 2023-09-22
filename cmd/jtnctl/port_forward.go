package jtnctl

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zakyyudha/jtnctl/config"
	"github.com/zakyyudha/jtnctl/pkg/kubectl"
	"os"
	"strings"
)

var portForwardCmd = &cobra.Command{
	Use:   "port-forward [service-name] [destination:source]",
	Short: "Forward ports for a Kubernetes service",
	Long:  "Forward ports for a Kubernetes service in the specified namespace. Use flags to configure the port forwarding.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.LoadConfig()
		if err != nil {
			fmt.Printf("Error reading config file")
			os.Exit(1)
		}

		namespace := config.Namespace
		serviceName := args[0]
		destinationSource := args[1]

		// Split the destination:source argument
		ports := strings.Split(destinationSource, ":")
		if len(ports) != 2 {
			fmt.Println("Invalid destination:source format. Use 'destination:source'.")
			os.Exit(1)
		}

		destinationPort := ports[0]
		sourcePort := ports[1]

		// Get the pod names matching the service name
		podToForward, err := kubectl.GetPodsName(namespace, serviceName)
		if err != nil {
			fmt.Printf("Error getting pod names: %v\n", err)
			os.Exit(1)
		}

		if len(podToForward) == 0 {
			fmt.Printf("No pods found matching '%s' in namespace '%s'\n", serviceName, namespace)
			return
		}

		// Execute the command using kubectl logs
		kubectlArgs := []string{"-n", namespace, "port-forward", podToForward, destinationPort + ":" + sourcePort}

		fmt.Println(kubectlArgs)

		// Run 'kubectl port-forward' on the selected pod
		if err := kubectl.RunCommand(kubectlArgs...); err != nil {
			fmt.Printf("Error executing 'kubectl port-forward': %v\n", err)
			os.Exit(1)
		}
	},
}
