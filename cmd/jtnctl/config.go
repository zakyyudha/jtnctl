package jtnctl

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zakyyudha/jtnctl/config"
	"strings"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings",
}

var (
	setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set configuration options",
		Args:  cobra.MinimumNArgs(1), // Require at least one key-value pair argument
		Run: func(cmd *cobra.Command, args []string) {
			// Parse and set configuration options
			for _, arg := range args {
				key, value, err := parseKeyValue(arg)
				if err != nil {
					fmt.Printf("Error parsing key-value pair: %v\n", err)
					return
				}
				// Set the configuration option
				if err := config.SetConfigOption(key, value); err != nil {
					fmt.Printf("Error setting configuration option: %v\n", err)
					return
				}
				fmt.Printf("Set configuration option: %s=%s\n", key, value)
			}
		},
	}

	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show configuration options",
		Run: func(cmd *cobra.Command, args []string) {
			// Get and display the configuration options
			configOptions := config.GetConfigOptions()
			for key, value := range configOptions {
				fmt.Printf("%s=%s\n", key, value)
			}
		},
	}
)

func init() {
	// Add subcommands to the "config" command
	configCmd.AddCommand(setCmd)
	configCmd.AddCommand(showCmd)
}

// Parse a key-value pair argument (e.g., "namespace=tds-stage")
func parseKeyValue(arg string) (string, string, error) {
	parts := strings.SplitN(arg, "=", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid key-value pair format: %s", arg)
	}
	return parts[0], parts[1], nil
}
