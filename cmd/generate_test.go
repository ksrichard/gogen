package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGenerateHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"generate", "--help"})
	rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	strOut := string(out)
	if !strings.Contains(strOut, generateCmd.Long) {
		t.Errorf("Output should contain '%s'!", generateCmd.Long)
	}
	if !strings.Contains(strOut, generateCmd.UsageString()) {
		t.Errorf("Output should contain '%s'!", generateCmd.UsageString())
	}
}

func TestGenerateRequiredFlags(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"generate"})
	cmdErr := rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	strOut := string(out)
	if !strings.Contains(strOut, generateCmd.UsageString()) {
		t.Errorf("Output should contain '%s'!", generateCmd.UsageString())
	}
	requiredFlagsStr := "required flag(s) \"output-dir\", \"template-dir\" not set"
	if cmdErr != nil && !strings.Contains(cmdErr.Error(), requiredFlagsStr) {
		t.Errorf("Output should contain '%s'!", requiredFlagsStr)
	}
}
