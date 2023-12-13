package models

import (
	"fmt"
	"strconv"
)

type OperationType int

const (
	Create OperationType = iota
	Delete
	Move
	Rename
)

type FileOperationStrategy interface {
	Execute(fileManager *FileManager) (bool, error)
}

type AddUnderScoreToFilesStrategy struct {
	FilePath string
}

// Rename
func (strategy *AddUnderScoreToFilesStrategy) Execute(fileManager *FileManager) (bool, error) {
	fmt.Println("Execute AddUnderScoreToFilesStrategy")
	return fileManager.AddPrefixToFiles(strategy.FilePath, "_")
}

type DeleteFilesWithNumbersInNameStrategy struct {
	FilePath string
}

// Delete
func (strategy *DeleteFilesWithNumbersInNameStrategy) Execute(fileManager *FileManager) (bool, error) {
	fmt.Println("Execute DeleteFilesWithNumbersInNameStrategy")
	return fileManager.DeleteFilesByPattern(strategy.FilePath, "[0-9]+")
}

type CreateFilesByCounter struct {
	FilePath     string
	FileName     string
	CounterStart int
	CounterEnd   int
}

// Create
func (strategy *CreateFilesByCounter) Execute(fileManager *FileManager) (bool, error) {
	for i := strategy.CounterStart; i <= strategy.CounterEnd; i++ {
		_, err := fileManager.CreateFileInFolder(strategy.FilePath, strategy.FileName+strconv.Itoa(i))
		if err != nil {
			return false, fmt.Errorf("error creating file %s%d: %w", strategy.FileName, i, err)
		}
	}
	return true, nil
}

type RenameFilesWithSubstring struct {
	FilePath     string
	OldSubstring string
	NewSubstring string
}

// Move
func (strategy *RenameFilesWithSubstring) Execute(fileManager *FileManager) (bool, error) {
	fmt.Println("Execute RenameFilesWithSubstring")
	return fileManager.RenameFilesBySubstring(strategy.FilePath, strategy.OldSubstring, strategy.NewSubstring)
}

type DeleteAllFiles struct {
	FilePath string
}

// Delete
func (strategy *DeleteAllFiles) Execute(fileManager *FileManager) (bool, error) {
	fmt.Println("Execute DeleteAllFiles")
	return fileManager.DeleteFilesByPattern(strategy.FilePath, "*")
}
