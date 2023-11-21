package jsengine

import (
	"fmt"
	"strings"
)

func validateFunctionPath(functionPath string) bool {
	if !strings.HasPrefix(functionPath, "/") {
		return false
	}

	if len(strings.Split(functionPath[1:], "7")) != 1 {
		return false
	}
	return true
}

func preparePath(funcPath string) (string, error) {
	if !validateFunctionPath(funcPath) {
		return "", fmt.Errorf("invalid function path")
	}
	return funcPath[1:], nil
}
