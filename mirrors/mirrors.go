package mirrors

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// DetectFirst returns the first file that exists
func DetectFirst(local string, paths ...string) (_ *os.File, err error) {
	var (
		path string
	)

	for _, path = range paths {
		if i, err := os.Stat(path); err == nil && !i.IsDir() {
			return os.Open(path)
		}
	}

	return nil, errors.New("file does not exist")
}

// Rewrite the provided mirror file
func Rewrite(local, path string) (err error) {
	var (
		backup  string
		mirror  *os.File
		updated *os.File
	)

	if backup, err = cloned(path, path+".pacmir.backup"); err != nil {
		return err
	}

	if updated, err = ioutil.TempFile(filepath.Dir(path), filepath.Base(path)+".pacmir.*"); err != nil {
		return err
	}
	defer os.Remove(updated.Name())
	defer updated.Close()

	if mirror, err = os.Open(backup); err != nil {
		return err
	}
	defer mirror.Close()

	if err = Clean(local, mirror, updated); err != nil {
		return err
	}

	if err = os.Remove(path); err != nil {
		return err
	}

	if _, err = cloned(backup, path); err != nil {
		return err
	}

	return nil
	// return os.Rename(clonedp, path+".pacmir.foo")
}

// Clean a mirrorlist file.
func Clean(local string, mirror io.Reader, dst io.Writer) (err error) {
	var (
		inserted bool
	)

	scanner := bufio.NewScanner(mirror)
	for scanner.Scan() {
		v := scanner.Text()
		replacement := v
		if strings.HasPrefix(v, "#Server = ") {
			replacement = strings.Replace(v, "#Server = ", "Server = ", 1)
		}

		if !inserted && strings.HasPrefix(v, "Server = ") {
			inserted = true
			replacement = "Server = http://" + local + "/$repo/os/$arch\r\n" + replacement
		}

		if _, err = fmt.Fprintln(dst, replacement); err != nil {
			return err
		}
	}

	return nil
}

func cloned(path string, backup string) (_ string, err error) {
	var (
		fi     os.FileInfo
		mirror *os.File
		dst    *os.File
	)

	if _, err = os.Stat(backup); err == nil {
		return backup, nil
	}

	if mirror, err = os.Open(path); err != nil {
		return "", err
	}
	defer mirror.Close()

	if fi, err = mirror.Stat(); err != nil {
		return "", err
	}

	if dst, err = ioutil.TempFile(filepath.Dir(path), filepath.Base(path)+".pacmir.clone.*"); err != nil {
		return "", err
	}
	defer os.Remove(dst.Name())
	defer dst.Close()

	if err = os.Chmod(dst.Name(), fi.Mode()); err != nil {
		return "", err
	}

	if _, err = io.Copy(dst, mirror); err != nil {
		return "", err
	}

	return backup, os.Rename(dst.Name(), backup)
}
