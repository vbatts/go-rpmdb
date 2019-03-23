package rpmdb

import (
	"bytes"
	"encoding/xml"
	"errors"
	"os/exec"
	"strings"
)

// Names returns the list of RPM packages installed.
//
// This is _just_ the simple name, which with kernels and multiarch can have
// duplicates. If you need unique keys, use NVRs().
func Names() ([]string, error) {
	cmd := exec.Command("rpm", "-qa", "--qf=%{NAME}\n")
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	err = isError(stderr, err)
	if err != nil {
		return nil, err
	}
	return strings.Fields(stdout.String()), nil
}

// NVRs returns the set unique name-version-release of packages installed.
func NVRs() ([]string, error) {
	cmd := exec.Command("rpm", "-qa")
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	err = isError(stderr, err)
	if err != nil {
		return nil, err
	}
	return strings.Fields(stdout.String()), nil
}

// Info collects detailed information on the unique package name (or NVR)
// provided.
func Info(nvr string) (interface{}, error) {
	cmd := exec.Command("rpm", "-q", "--xml", nvr)
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	err = isError(stderr, err)
	if err != nil {
		return nil, err
	}

	hdr := RPMHeader{}
	// parse dat XML
	err = xml.Unmarshal(stdout.Bytes(), &hdr)
	if err != nil {
		return nil, err
	}

	return hdr, nil
}

// isError parses the output of `rpm` since sometimes (all the time?) errors
// are shown, but 0 is still returned
func isError(stderr *bytes.Buffer, err error) error {
	if err != nil {
		return err
	}
	if strings.Contains(strings.ToLower(stderr.String()), "error") {
		return errors.New(stderr.String())
	}
	return nil
}
