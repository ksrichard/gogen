package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestRootExecute(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{})
	rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	strOut := string(out)
	if !strings.Contains(strOut, rootCmd.Long) {
		t.Errorf("Output should contain '%s'!", rootCmd.Long)
	}
	if !strings.Contains(strOut, rootCmd.UsageString()) {
		t.Errorf("Output should contain '%s'!", rootCmd.UsageString())
	}
}

func TestRootExecuteHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"--help"})
	rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	strOut := string(out)
	if !strings.Contains(strOut, rootCmd.Long) {
		t.Errorf("Output should contain '%s'!", rootCmd.Long)
	}
	if !strings.Contains(strOut, rootCmd.UsageString()) {
		t.Errorf("Output should contain '%s'!", rootCmd.UsageString())
	}
}

func TestRootExecuteHelpShort(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"-h"})
	rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	strOut := string(out)
	if !strings.Contains(strOut, rootCmd.Long) {
		t.Errorf("Output should contain '%s'!", rootCmd.Long)
	}
	if !strings.Contains(strOut, rootCmd.UsageString()) {
		t.Errorf("Output should contain '%s'!", rootCmd.UsageString())
	}
}

func TestRootExecuteHelpCmd(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"help"})
	rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	strOut := string(out)
	if !strings.Contains(strOut, rootCmd.Long) {
		t.Errorf("Output should contain '%s'!", rootCmd.Long)
	}
	if !strings.Contains(strOut, rootCmd.UsageString()) {
		t.Errorf("Output should contain '%s'!", rootCmd.UsageString())
	}
}
