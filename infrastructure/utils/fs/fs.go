package fs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

func PathExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func CreateDir(path string, mode os.FileMode) error {
	return os.MkdirAll(path, mode)
}

func CreateDirs(mode os.FileMode, path ...string) error {
	for _, p := range path {
		if err := CreateDir(p, mode); err != nil {
			return err
		}
	}

	return nil
}

func Remove(path string) error {
	return os.RemoveAll(path)
}

func BatchRemove(path ...string) error {
	for _, p := range path {
		if err := Remove(p); err != nil {
			return err
		}
	}

	return nil
}

func CopyRecursively(srcPath, dstPath string) error {
	if !PathExists(dstPath) {
		if err := CreateDir(dstPath, os.ModePerm); err != nil {
			return err
		}
	}

	if IsFile(srcPath) {
		if err := CopyFile(srcPath, dstPath); err != nil {
			return err
		}
		fmt.Println(111)
		return nil
	}

	return filepath.Walk(srcPath, func(fname string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcPath, fname)
		if err != nil {
			return err
		}

		mirrorPath := filepath.Join(dstPath, relPath)

		if info.IsDir() {
			if err := CreateDir(mirrorPath, info.Mode()); err != nil {
				return err
			}
		} else {
			fmt.Println(fname, mirrorPath)
			if err := CopyFile(fname, mirrorPath); err != nil {
				return err
			}
		}

		return nil
	})
}

func CopyFile(srcPath, dstPath string) error {
	stat, err := os.Lstat(srcPath)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return nil
	}

	dstFile := dstPath
	if IsDir(dstPath) {
		dstFile = filepath.Join(dstPath, stat.Name())
	}

	dstf, err := os.OpenFile(dstFile, os.O_CREATE|os.O_RDWR, stat.Mode())
	if err != nil {
		return err
	}
	defer dstf.Close()

	srcf, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcf.Close()

	_, err = io.Copy(dstf, srcf)
	if err != nil {
		return err
	}

	return nil
}
