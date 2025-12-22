package model

import "time"

type TestStatus string

const (
	TestPass TestStatus = "PASSSED"
	TestFail TestStatus = "FAILED"
)

type TestOpts struct {
	CoverageFlag bool
}

type TestResults struct {
	Status   TestStatus
	Duration time.Duration
	Failure  *FailureCapture
	Stdout   string
	Stderr   string
}

type FailureCapture struct {
	File    string
	Line    int
	Message string
}
