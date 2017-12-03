package libtree

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Run executes the main tree builder
func Run(m string) {
	var p string
	if runtime.GOOS == "windows" {
		p = "C:\\"
	} else {
		p = "/"
	}

	fmt.Println(fmt.Sprintf("[+] Mode: %s", m))
	if m == "directory" {
		filepath.Walk(p, WalkerDir)
	} else {
		filepath.Walk(p, WalkerFile)
	}
}

func WalkerDir(path string, info os.FileInfo, err error) error {
	e := Walker(true, path, info, err)
	return e
}

func WalkerFile(path string, info os.FileInfo, err error) error {
	e := Walker(false, path, info, err)
	return e
}

func Walker(dirOnly bool, path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(fmt.Sprintf("[!] %s", err.Error()))
		return nil
	}

	l1 := len(strings.Split(path, "/"))
	l2 := len(strings.Split(path, "\\"))

	var length int
	if l1 > l2 {
		length = l1
	} else {
		length = l2
	}

	s := fmt.Sprintf("%s└──", strings.Repeat("  ", length))

	if dirOnly {
		if info.IsDir() {
			fmt.Println(fmt.Sprintf(
				"%s %s [Size: %d]",
				s,
				info.Name(),
				info.Size()),
			)
		}
	} else {
		var t string
		if info.IsDir() {
			t = "dir"
		} else {
			t = "file"
		}
		fmt.Println(fmt.Sprintf(
			"%s %s [Size: %d | Type: %s]",
			s,
			info.Name(),
			info.Size(),
			t),
		)
	}
	return nil
}
