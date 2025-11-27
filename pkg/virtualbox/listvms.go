package virtualbox

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func ListVMs(ctx context.Context) ([]*VM, error) {
	cmd := exec.CommandContext(ctx, "vboxmanage", "list", "vms")

	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		stderrMsg := strings.TrimSpace(stderr.String())
		if stderrMsg != "" {
			return nil, fmt.Errorf("%w: %w: %s", ErrUnknownError, err, stderrMsg)
		}

		return nil, fmt.Errorf("%w: %w", ErrUnknownError, err)
	}

	return parseVMs(stdout.String()), nil
}

func parseVMs(output string) []*VM {
	raw := strings.TrimSpace(output)
	if raw == "" {
		return nil
	}

	lines := strings.Split(raw, "\n")

	var vms []*VM

	for _, line := range lines {
		name, uuid, ok := parseVMLine(line)
		if !ok {
			continue
		}

		vms = append(vms, NewVM(uuid, name))
	}

	return vms
}

func parseVMLine(line string) (string, string, bool) {
	line = strings.TrimSpace(line)
	if line == "" {
		return "", "", false
	}

	nameStart := strings.LastIndex(line, " ")
	if nameStart == -1 {
		return "", "", false
	}

	name := line[1 : nameStart-1]
	uuid := line[nameStart+2 : len(line)-1]

	return name, uuid, true
}
