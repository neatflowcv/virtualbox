package virtualbox

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func StartVM(ctx context.Context, uuid string) error {
	cmd := exec.CommandContext(ctx, "vboxmanage", "startvm", "--type", "headless", uuid)

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

		if strings.Contains(stderr.String(), "is already locked by a session") {
			return ErrVMAlreadyLocked
		}

		stderrMsg := strings.TrimSpace(stderr.String())
		if stderrMsg != "" {
			return fmt.Errorf("%w: %w: %s", ErrUnknownError, err, stderrMsg)
		}

		return fmt.Errorf("%w: %w", ErrUnknownError, err)
	}

	return nil
}
