package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"hash"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var xJar = XJar{
	md5:  []byte{#{xJar.md5}},
	sha1: []byte{#{xJar.sha1}},
}

var xKey = XKey{
	algorithm: []byte{#{xKey.algorithm}},
	keysize:   []byte{#{xKey.keysize}},
	ivsize:    []byte{#{xKey.ivsize}},
	password:  []byte{#{xKey.password}},
}

var xJava = XJava{
	javaversion:  []byte{#{xJava.javaversion}},
	javamd5:      []byte{#{xJava.javamd5}},
	javasha1:     []byte{#{xJava.javasha1}},
}

func main() {
	// search the jar to start
	jar, err := JAR(os.Args)
	if err != nil {
		panic(err)
	}

	// parse jar name to absolute path
	jarPath, err := filepath.Abs(jar)
	if err != nil {
		panic(err)
	}

	// verify jar with MD5
	jarMD5, err := MD5(jarPath)
	if err != nil {
		panic(err)
	}
	if bytes.Compare(jarMD5, xJar.md5) != 0 {
		panic(errors.New("invalid jar with MD5"))
	}

	// verify jar with SHA-1
	jarSHA1, err := SHA1(jarPath)
	if err != nil {
		panic(err)
	}
	if bytes.Compare(jarSHA1, xJar.sha1) != 0 {
		panic(errors.New("invalid jar with SHA-1"))
	}

	// check agent forbid
	{
		args := os.Args
		l := len(args)
		for i := 0; i < l; i++ {
			arg := args[i]
			if strings.HasPrefix(arg, "-javaagent:") {
				panic(errors.New("agent forbidden"))
			}
		}
	}

	// start java application
	java := os.Args[1]
	args := os.Args[2:]

    // first cmd must java
    if !strings.EqualFold(java, "java") && !strings.HasSuffix(java, "/java") && !strings.HasSuffix(java, "\\java") && !strings.HasSuffix(java, "/java.exe") && !strings.HasSuffix(java, "\\java.exe") {
        panic(errors.New("not support cmd, only support java"))
    }

    javaversion := string(xJava.javaversion)

    // 检查JDK版本
    if err := checkJDKVersion(java, javaversion); err != nil {
    	panic(err)
    }

    // 检查JDK MD5
    if err := checkJDKMD5(java, xJava.javamd5); err != nil {
    	panic(err)
    }

    // 检查JDK SHA1
    if err := checkJDKSHA1(java, xJava.javasha1); err != nil {
    	panic(err)
    }

    // 添加java参数 -XX:+DisableAttachMechanism
    args = append([]string{"-XX:+DisableAttachMechanism"},args...)

	key := bytes.Join([][]byte{
		xKey.algorithm, {13, 10},
		xKey.keysize, {13, 10},
		xKey.ivsize, {13, 10},
		xKey.password, {13, 10},
	}, []byte{})
	cmd := exec.Command(java, args...)
	cmd.Stdin = bytes.NewReader(key)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

// find jar name from args
func JAR(args []string) (string, error) {
	var jar string

	l := len(args)
	for i := 1; i < l-1; i++ {
		arg := args[i]
		if arg == "-jar" {
			jar = args[i+1]
		}
	}

	if jar == "" {
		return "", errors.New("unspecified jar name")
	}

	return jar, nil
}

// calculate file's MD5
func MD5(path string) ([]byte, error) {
	return HASH(path, md5.New())
}

// calculate file's SHA-1
func SHA1(path string) ([]byte, error) {
	return HASH(path, sha1.New())
}

// calculate file's HASH value with specified HASH Algorithm
func HASH(path string, hash hash.Hash) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	_, _err := io.Copy(hash, file)
	if _err != nil {
		return nil, _err
	}

	sum := hash.Sum(nil)

	return sum, nil
}

// 检查JDK版本
func checkJDKVersion(java string, requiredVersion string) error {
	cmd := exec.Command(java, "-version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	versionStr := string(output)

	if !strings.EqualFold(requiredVersion, "unknown") && !strings.Contains(versionStr, requiredVersion) {
		return errors.New("JDK version does not meet the requirement")
	}
	return nil
}

// 检查JDK MD5
func checkJDKMD5(java string, requiredMD5 []byte) error {
    // 获取Java执行器路径
    javaExecutable, err := exec.LookPath(java)
    if err != nil {
        return err
    }

    // 计算Java执行器的MD5
    javaMD5, err := MD5(javaExecutable)
    if err != nil {
        return err
    }

    md5Str := string(requiredMD5)

    // 验证Java执行器的MD5
    if !strings.EqualFold(md5Str, "unknown") && bytes.Compare(javaMD5, requiredMD5) != 0 {
        return errors.New("Java executable has been tampered with")
    }
    return nil
}

// 检查JDK SHA1
func checkJDKSHA1(java string, requiredSHA1 []byte) error {
    // 获取Java执行器路径
    javaExecutable, err := exec.LookPath(java)
    if err != nil {
        return err
    }

    // 计算Java执行器的SHA1
    javaSHA1, err := SHA1(javaExecutable)
    if err != nil {
        return err
    }

    sha1Str := string(requiredSHA1)

    // 验证Java执行器的SHA1
    if !strings.EqualFold(sha1Str, "unknown") && bytes.Compare(javaSHA1, requiredSHA1) != 0 {
        return errors.New("Java executable has been tampered with")
    }
    return nil
}

type XJar struct {
	md5  []byte
	sha1 []byte
}

type XKey struct {
	algorithm []byte
	keysize   []byte
	ivsize    []byte
	password  []byte
}

type XJava struct {
	javaversion []byte
	javamd5     []byte
	javasha1    []byte
}
