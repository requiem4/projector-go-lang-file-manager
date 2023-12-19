package models

import (
	"strconv"
)

type FileOperationStrategy interface {
	Execute(fileManager *FileManager) error
}

type AddUnderScoreToFilesStrategy struct {
	FilePath string
}

// Rename
func (strategy *AddUnderScoreToFilesStrategy) Execute(fileManager *FileManager) error {
	if fileManager.logger != nil {
		fileManager.logger.Printf("Execute AddUnderScoreToFilesStrategy")
	}
	return fileManager.AddPrefixToFiles(strategy.FilePath, "_")
}

type DeleteFilesWithNumbersInNameStrategy struct {
	FilePath string
}

// Delete
func (strategy *DeleteFilesWithNumbersInNameStrategy) Execute(fileManager *FileManager) error {
	if fileManager.logger != nil {
		fileManager.logger.Printf("Execute DeleteFilesWithNumbersInNameStrategy")
	}
	return fileManager.DeleteFilesByPattern(strategy.FilePath, "[0-9]+")
}

type CreateFilesByCounter struct {
	FilePath     string
	FileName     string
	CounterStart int
	CounterEnd   int
}

// Create
func (strategy *CreateFilesByCounter) Execute(fileManager *FileManager) error {
	if fileManager.logger != nil {
		fileManager.logger.Printf("Execute CreateFilesByCounter")
	}
	for i := strategy.CounterStart; i <= strategy.CounterEnd; i++ {
		err := fileManager.CreateFileInFolder(strategy.FilePath, strategy.FileName+strconv.Itoa(i))
		if err != nil {
			if fileManager.logger != nil {
				fileManager.logger.Printf("error creating file %s%d: %w", strategy.FileName, i, err)
			}
			return err
		}
	}
	if fileManager.logger != nil {
		fileManager.logger.Printf("CreateFilesByCounter success")
	}
	return nil
}

type RenameFilesWithSubstring struct {
	FilePath     string
	OldSubstring string
	NewSubstring string
}

// Move
func (strategy *RenameFilesWithSubstring) Execute(fileManager *FileManager) error {
	if fileManager.logger != nil {
		fileManager.logger.Printf("Execute RenameFilesWithSubstring")
	}
	return fileManager.RenameFilesBySubstring(strategy.FilePath, strategy.OldSubstring, strategy.NewSubstring)
}

type DeleteAllFiles struct {
	FilePath string
}

// Delete
func (strategy *DeleteAllFiles) Execute(fileManager *FileManager) error {
	if fileManager.logger != nil {
		fileManager.logger.Printf("Execute DeleteAllFiles")
	}
	return fileManager.DeleteFilesByPattern(strategy.FilePath, ".*")
}
