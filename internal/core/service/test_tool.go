package service

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gunzgo2mars/achi-cli/internal/core/model"
)

func (s *promptService) RunAllTests(opts *model.TestOpts) *model.TestResults {

	pkgs := []string{"./..."}
	args := []string{"test"}

	if opts.CoverageFlag {
		args = append(args, "-v")
		args = append(args, "-coverprofile=coverage.out")
	}

	args = append(args, pkgs...)

	cmd := exec.Command("go", args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	start := time.Now()
	err := cmd.Run()
	duration := time.Since(start)

	result := &model.TestResults{
		Duration: duration,
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		Status:   model.TestPass,
	}

	if err != nil {
		result.Status = model.TestFail
		result.Failure = extractFailure(result.Stdout + result.Stderr)
	}

	return result
}

func (s *promptService) GetTotalCoverage() (float64, error) {
	cmd := exec.Command("go", "tool", "cover", "-func=coverage.out")

	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	fmt.Printf("Out: %s \n", string(out))

	for line := range bytes.SplitSeq(out, []byte("\n")) {
		s := string(line)
		if strings.HasPrefix(s, "total:") {
			fields := strings.Fields(s)
			percent := strings.TrimSuffix(fields[len(fields)-1], "%")
			return strconv.ParseFloat(percent, 64)
		}
	}

	return 0, nil
}

func extractFailure(output string) *model.FailureCapture {

	var failureRegex = regexp.MustCompile(`(?m)^(.+_test\.go):(\d+):\s+(.*)$`)
	matches := failureRegex.FindStringSubmatch(output)
	if len(matches) != 4 {
		return nil
	}

	line, _ := strconv.Atoi(matches[2])

	return &model.FailureCapture{
		File:    matches[1],
		Line:    line,
		Message: matches[3],
	}
}
