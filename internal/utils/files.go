package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type File struct{
	FilePath   []string
	ConfigPath []string
}

//creates file based on filePath and writes data from config files, when no config file needed, then pass empty string arr
func CreateFiles(filePath []string, configPath []string, projectName string){
	path := filepath.Join(filePath...)
	file,_ := os.Create(path)

	if(configPath != nil){
		content, err := getConfigFileContent(projectName, configPath...)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	
		writeFileContent(file, content)
	}

	defer file.Close()
}

func getConfigFileContent(projectName string, filePath ...string) (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	executableDir := filepath.Dir(executablePath)

	paths := append([]string{executableDir}, filePath...)
	configPath := filepath.Join(paths...)

	content, err := os.ReadFile(configPath)
	if err != nil {
		return "", err
	}

	replacedContent := strings.ReplaceAll(string(content), "XPROJECTNAMEX", projectName)

	return replacedContent, nil
}

func writeFileContent(file *os.File, content string) {
	_, err := file.WriteString(content)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}