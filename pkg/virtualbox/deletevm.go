package virtualbox

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func DeleteVM(ctx context.Context, uuid string) error {
	cmd := exec.CommandContext(ctx, "vboxmanage", "unregistervm", "--delete-all", uuid)

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

		stderrMsg := strings.TrimSpace(stderr.String())
		if stderrMsg != "" {
			return fmt.Errorf("%w: %w: %s", ErrUnknownError, err, stderrMsg)
		}

		return fmt.Errorf("%w: %w", ErrUnknownError, err)
	}

	return nil
}
