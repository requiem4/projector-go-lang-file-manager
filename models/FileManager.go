package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileManager struct {
}

func (fileManager FileManager) CreateFileInFolder(folderPath string, fileName string) (bool, error) {

	filePath := filepath.Join(folderPath, fileName)

	_, err := os.Stat(filePath)

	if err == nil {
		fmt.Printf("File %s exists in folder %s.\n", fileName, folderPath)
		return true, nil
	}

	if os.IsNotExist(err) {
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return false, err
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()
	fmt.Printf("File %s has been created successfully.\n", fileName)
	return true, nil
}

func (fileManager FileManager) CreateFolder(folderPath string) (bool, error) {
	if err := os.MkdirAll(folderPath, 0755); err != nil {
		fmt.Printf("Error creating nested directories %s: %s\n", folderPath, err)
		return false, err
	} else {
		fmt.Printf("Directory %s created successfully\n", folderPath)
		return true, nil
	}
}

func (fileManager FileManager) DeleteFile(folderPath, fileName string) (bool, error) {

	filePath := filepath.Join(folderPath, fileName)

	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("Error deleting file %s: %s\n", fileName, err)
		return false, err
	} else {
		fmt.Printf("File %s deleted successfully.\n", fileName)
		return true, nil
	}
}

func (fileManager FileManager) DeleteFileAndFolder(folderPath, fileName string) (bool, error) {

	_, err := fileManager.DeleteFile(folderPath, fileName)

	err = os.Remove(folderPath)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Printf("Folder %s deleted successfully.\n", folderPath)
	return true, nil
}

func (fileManager FileManager) CreateFilesList(folderPath string, fileNames []string) bool {
	allFilesCreated := true
	for _, fileName := range fileNames {
		_, err := fileManager.CreateFileInFolder(folderPath, fileName)
		if err != nil {
			allFilesCreated = false
		}
	}
	return allFilesCreated
}

func (fileManager FileManager) DeleteFilesList(folderPath string, fileNames []string) bool {
	allFilesDeleted := true
	for _, fileName := range fileNames {
		_, err := fileManager.DeleteFile(folderPath, fileName)
		if err != nil {
			allFilesDeleted = false
		}
	}
	return allFilesDeleted
}

func (fileManager FileManager) DeleteFilesBySubstring(basePath string, substring string) error {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return err
	}

	for _, f := range files {
		if strings.Contains(f.Name(), substring) {
			err := os.RemoveAll(filepath.Join(basePath, f.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (fileManager FileManager) RenameFilesBySubstring(basePath string, oldSubstring string, newSubstring string) (bool, error) {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return false, err
	}

	for _, f := range files {
		if strings.Contains(f.Name(), oldSubstring) {
			fileName := f.Name()
			oldPath := filepath.Join(basePath, fileName)
			newName := strings.ReplaceAll(fileName, oldSubstring, newSubstring)

			newPath := filepath.Join(basePath, newName)

			err := os.Rename(oldPath, newPath)
			if err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

func (fileManager FileManager) DeleteFilesByPattern(basePath string, regex string) (bool, error) {
	compiledRegex, err := regexp.Compile(regex)
	if err != nil {
		fmt.Errorf("invalid pattern: %w", err)
		return false, err
	}

	files, err := ioutil.ReadDir(basePath)

	if err != nil {
		return false, err
	}

	for _, file := range files {
		if compiledRegex.MatchString(file.Name()) {
			err := os.RemoveAll(filepath.Join(basePath, file.Name()))
			if err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

func (fileManager FileManager) AddPrefixToFiles(basePath, prefix string) (bool, error) {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		oldPath := filepath.Join(basePath, file.Name())
		newPath := filepath.Join(basePath, prefix+file.Name())
		err := os.Rename(oldPath, newPath)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (fileManager FileManager) CopyFilesToNewDir(srcPath, destPath string) error {
	files, err := ioutil.ReadDir(srcPath)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(destPath, 0755); err != nil {
		return err
	}

	for _, f := range files {
		srcFile := filepath.Join(srcPath, f.Name())
		destFile := filepath.Join(destPath, f.Name())
		_, err := fileManager.CopyFileContent(srcFile, destFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fileManager FileManager) CopyFileContent(sourceFile string, destinationFile string) (bool, error) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	err = ioutil.WriteFile(destinationFile, input, 0755)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
