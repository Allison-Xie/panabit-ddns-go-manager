package env

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func CopyFile(src, dest string, perm os.FileMode) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()
	d, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer d.Close()
	if _, err := io.Copy(d, s); err != nil {
		return err
	}
	if err := d.Chmod(perm); err != nil {
		return err
	}
	return nil
}

func CopyDir(src, dest string, perm os.FileMode) error {
	if err := os.MkdirAll(dest, perm); err != nil {
		return err
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sp := filepath.Join(src, entry.Name())
		dp := filepath.Join(dest, entry.Name())
		if entry.IsDir() {
			if err := CopyDir(sp, dp, perm); err != nil {
				return err
			}
		} else {
			if err := CopyFile(sp, dp, perm); err != nil {
				return err
			}
		}
	}
	return nil
}

func ReadPidfile(name string) (int, error) {
	d, err := os.ReadFile(name)
	if err != nil {
		return 0, err
	}
	if pid, err := strconv.Atoi(strings.TrimSpace(string(d))); err != nil {
		return 0, err
	} else {
		return pid, nil
	}
}

func WritePidfile(name string, pid int) error {
	return os.WriteFile(name, []byte(strconv.Itoa(pid)), 0644)
}

func RemovePidfile(name string) error {
	return os.Remove(name)
}

func DescribeProcessExists(pid int) (bool, error) {
	p, err := os.FindProcess(pid)
	if err != nil {
		return false, err
	}
	return p.Signal(syscall.Signal(0)) == nil, nil
}
