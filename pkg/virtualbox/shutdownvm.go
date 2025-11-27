package virtualbox

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func ShutdownVM(ctx context.Context, uuid string) error {
	cmd := exec.CommandContext(ctx, "vboxmanage", "controlvm", uuid, "shutdown")

	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if strings.Contains(stderr.String(), "Could not find a registered machine named") {
			return ErrVMNotFound
		}

		if strings.Contains(stderr.String(), "is not currently running") {
			return ErrVMNotRunning
		}

		stderrMsg := strings.TrimSpace(stderr.String())
		if stderrMsg != "" {
			return fmt.Errorf("%w: %w: %s", ErrUnknownError, err, stderrMsg)
		}

		return fmt.Errorf("%w: %w", ErrUnknownError, err)
	}

	return nil
}
