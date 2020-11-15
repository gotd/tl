package main

import (
	"bufio"
	"bytes"
	"crypto/md5" // #nosec
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := filepath.Walk("_testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		data, readErr := ioutil.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		s := bufio.NewScanner(bytes.NewReader(data))
		for s.Scan() {
			text := s.Text()
			// name, O_RDWR|O_CREATE|O_TRUNC, 0666
			if !strings.HasSuffix(text, ";") {
				continue
			}
			crc := md5.Sum(s.Bytes())// #nosec G401

			targetName := fmt.Sprintf("testdata-%x", crc)
			targetPath := filepath.Join("_fuzz", "definitions", "corpus", targetName)

			if err := ioutil.WriteFile(targetPath, []byte(strings.TrimSuffix(text, ";")), 0600); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		panic(err)
	}
}
