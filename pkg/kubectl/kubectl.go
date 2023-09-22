package kubectl

import (
	"fmt"
	"github.com/zakyyudha/jtnctl/shared"
	"os"
	"os/exec"
	"strings"
)

// GetPodsName returns the names of pods matching a search string in the specified namespace.
func GetPodsName(namespace, search string) (string, error) {
	cmd := exec.Command("kubectl", "-n", namespace, "get", "pods", "--no-headers", "-o", "custom-columns=:metadata.name")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	podNames := strings.Fields(string(output))
	findPod := shared.FindSubstring(podNames, search)
	fmt.Println("Pods name ->", findPod)
	return findPod, nil
}

// RunCommand runs a kubectl command with the given arguments.
func RunCommand(args ...string) error {
	cmd := exec.Command("kubectl", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running kubectl command: %v", err)
	}

	return nil
}
