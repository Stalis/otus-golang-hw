package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func FindVar(envVars []string, name string) int {
	for i, v := range envVars {
		if curName, _, _ := strings.Cut(v, "="); curName == name {
			return i
		}
	}

	return -1
}

func RemoveItem(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...)

	keys := make([]string, 0, len(env))
	for k := range env {
		keys = append(keys, k)
		if i := FindVar(command.Env, k); i > 0 {
			command.Env = RemoveItem(command.Env, i)
		}
	}

	for k, v := range env {
		if v.NeedRemove {
			continue
		}

		envVar := fmt.Sprintf("%v=%v", k, v.Value)
		command.Env = append(command.Env, envVar)
	}

	command.Run()

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.ProcessState.ExitCode()
}
