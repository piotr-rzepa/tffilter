package terraform

import (
	"fmt"
	"log"
	"os/exec"
	"slices"
)

func GetBinary(binary string) string {
	binaries := []string{"terraform", "terragrunt"}
	if !slices.Contains(binaries, binary) {
		log.Fatalf("'%s' binary is not supported", binary)
	}
	path, err := exec.LookPath(binary)
	if err != nil {
		log.Fatalf("'%s' binary not found in $PATH\n", binary)
	}
	fmt.Printf("'%s' binary available at '%s'\n", binary, path)
	return path
}
