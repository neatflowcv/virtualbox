package virtualbox

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func ShowVMInfo(ctx context.Context, uuid string) (*VMInfo, error) {
	cmd := exec.CommandContext(ctx, "vboxmanage", "showvminfo", "--machinereadable", uuid)

	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if strings.Contains(stderr.String(), "Could not find a registered machine named") {
			return nil, ErrVMNotFound
		}

		stderrMsg := strings.TrimSpace(stderr.String())
		if stderrMsg != "" {
			return nil, fmt.Errorf("%w: %w: %s", ErrUnknownError, err, stderrMsg)
		}

		return nil, fmt.Errorf("%w: %w", ErrUnknownError, err)
	}

	lines := strings.Split(stdout.String(), "\n")
	name := ""
	status := ""

	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "name="):
			name = strings.TrimPrefix(line, "name=")
			name = name[1 : len(name)-1]
		case strings.HasPrefix(line, "VMState="):
			status = strings.TrimPrefix(line, "VMState=")
			status = status[1 : len(status)-1]
		}
	}

	return NewVMInfo(uuid, name, status), nil
}
