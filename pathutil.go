// Package pathutil contains utility functions for working paths
// It is also a wrapper for path & filepath which can be accessed directly
package pathutil

import (
	"os"
	"os/user"
	pathlib "path"
)

// Check if a path exists regardless of whether it is a dir or file
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Expands '~/' prefix to the home directory
func Expand(path string) string {
	if len(path) < 2 {
		return path
	}
	if path[:2] == "~/" {
		usr, err := user.Current()
        if err != nil {
            panic(err)
        }
		home := usr.HomeDir
		path = pathlib.Join(home, path[2:])
	}
	return path
}
