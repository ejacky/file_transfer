package core

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

const (
	// FileSha1CMD : 计算文件sha1值
	FileSha1CMD = `
	#!/bin/bash
	sha1sum $1 | awk '{print $1}'
	`
)

func ComputeSha1ByShell(destPath string) (string, error) {
	cmdStr := strings.Replace(FileSha1CMD, "$1", destPath, 1)
	hashCmd := exec.Command("bash", "-c", cmdStr)
	if filehash, err := hashCmd.Output(); err != nil {
		fmt.Println(err)
		return "", err
	} else {
		reg := regexp.MustCompile("\\s+")
		return reg.ReplaceAllString(string(filehash), ""), nil
	}
}

func TestCalcSha1(t *testing.T) {
	fmt.Println(calcFileSha1("../server/16eab20f299f57b8/2"))
	fmt.Println(calcChunkSha1("../server/16eab20f299f57b8/2"))
	fmt.Println("hello world!")
	fmt.Println(ComputeSha1ByShell("../server/16eab20f299f57b8/2"))

}
