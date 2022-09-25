package main

import "strings"

type Job struct {
	label  string
	output string
	err    error
}

func NewJob(label string) *Job {
	return &Job{
		label: label,
	}
}

func (j *Job) Complete(output string) {
	j.output = output
}

func (j *Job) Fail(err error) {
	j.err = err
}

func (j *Job) String() string {
	var sb strings.Builder
	sb.WriteString(j.label)
	sb.WriteString("... ")
	if j.err != nil {
		sb.WriteString("Failed: ")
		sb.WriteString(j.err.Error())
	} else if j.output != "" {
		sb.WriteString("Done: ")
		sb.WriteString(j.output)
	}
	return sb.String()
}
