package function_module

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
)

type FunctionModule struct {
	Name     string
	Language string
	Source   string
	Method   string
	Path     string
	Cpu      string
	Memory   string
}

func (F *FunctionModule) Build() (err error) {
	fileData, err := base64.StdEncoding.DecodeString(F.Source)
	if err != nil {
		return err
	}
	path := "data"
	if _, err = os.Stat(path); err == nil {
		fmt.Println("path exists 1", path)
	} else {
		fmt.Println("path not exists ", path)
		err = os.MkdirAll(path, os.ModePerm)

		if err != nil {
			fmt.Println("Error creating directory")
			fmt.Println(err)
			return err
		}
	}
	file, err1 := os.Create("data/source.tar.gz")
	if err1 != nil {
		return err1
	}
	file.Write(fileData)

	dockerFile := fmt.Sprint("templates/", F.Language, ".dockerfile")
	fmt.Println(dockerFile)
	funcName := fmt.Sprint("function-", F.Name)
	cmd := exec.Command("docker", "build", "-f", dockerFile, "-t", funcName, "data/")
	err2 := cmd.Run()
	if err2 != nil {
		return err2
	}
	os.Remove("data/source.tar.gz")
	return nil
}

func (F *FunctionModule) Run() (resp string, err error) {
	cpuConfig := fmt.Sprint("--cpus=", F.Cpu)
	memConfig := fmt.Sprint("--memory=", F.Memory)
	funcName := fmt.Sprint("function-", F.Name)
	cmd := exec.Command("docker", "run", "--rm", cpuConfig, memConfig, funcName)
	out, err := cmd.CombinedOutput()
	//
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (F *FunctionModule) Delete() (err error) {
	funcName := fmt.Sprint("function-", F.Name)
	cmd := exec.Command("docker", "rmi", funcName)
	err = cmd.Run()
	return err
}
