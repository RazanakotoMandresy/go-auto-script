package main

import (
	"slices"
	"testing"
)

func TestScript(t *testing.T) {
	t.Run("ReadFile", func(t *testing.T) {
		want := []string{"ls", "ls --help", "docker version"}
		got, _ := arrScript("../script_test.txt")

		if !slices.Equal(want, got) {
			t.Errorf("want %v but got %v", want, got)
		}
	})
	t.Run("command executions", func(t *testing.T) {
		arrCmd, _ := arrScript("../script_test.txt")
		if err := cmdExecute(arrCmd); err != nil {
			t.Errorf("cmdExecute error %v", err)
		}
	})

}
