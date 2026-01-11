package lib

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type ToolRequirement struct {
	Name          string
	VersionArgs   []string
	MinVersion    string
	VersionRegexp *regexp.Regexp
}

var semverRegexp = regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`)

func CheckTools(tools []ToolRequirement) error {
	var errors []string

	for _, tool := range tools {
		// Presence check
		if _, err := exec.LookPath(tool.Name); err != nil {
			errors = append(errors,
				fmt.Sprintf("%s is not installed or not in PATH", tool.Name),
			)
			continue
		}

		// Version check
		if tool.MinVersion != "" {
			if err := checkVersion(tool); err != nil {
				errors = append(errors,
					fmt.Sprintf("%s: %s", tool.Name, err.Error()),
				)
			}
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf(
			"tooling requirements not met:\n  - %s",
			strings.Join(errors, "\n  - "),
		)
	}

	return nil
}

func checkVersion(tool ToolRequirement) error {
	cmd := exec.Command(tool.Name, tool.VersionArgs...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to read version")
	}

	re := tool.VersionRegexp
	if re == nil {
		re = semverRegexp
	}

	matches := re.FindStringSubmatch(out.String())
	if len(matches) < 4 {
		return fmt.Errorf("unable to parse version output")
	}

	current := matches[1:4]
	required := parseVersion(tool.MinVersion)

	if compareVersions(current, required) < 0 {
		return fmt.Errorf(
			"version %s+ required (found %s)",
			tool.MinVersion,
			strings.Join(current, "."),
		)
	}

	return nil
}

func parseVersion(v string) []string {
	m := semverRegexp.FindStringSubmatch(v)
	if len(m) < 4 {
		return []string{"0", "0", "0"}
	}
	return m[1:4]
}

func compareVersions(a, b []string) int {
	for i := 0; i < 3; i++ {
		ai, _ := strconv.Atoi(a[i])
		bi, _ := strconv.Atoi(b[i])

		if ai < bi {
			return -1
		}
		if ai > bi {
			return 1
		}
	}
	return 0
}
