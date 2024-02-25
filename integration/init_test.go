package integration

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var (
	builtBinaryPath string
)

func TestMain(m *testing.M) {
	tmpDir, err := os.MkdirTemp("", "integration")
	if err != nil {
		panic("failed to create temp dir")
	}
	defer os.RemoveAll(tmpDir)

	builtBinaryPath = filepath.Join(tmpDir, "TestCalculator")

	cmd := exec.Command("go", "build", "-o", builtBinaryPath, "..")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("failed to build TestCalculator: %s", output))
	}
	code := m.Run()
	os.Exit(code)
}
