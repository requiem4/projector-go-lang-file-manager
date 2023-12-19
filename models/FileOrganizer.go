package models

import "log"

type FileOrganizer struct {
	Strategy FileOperationStrategy
	Logger   *log.Logger
}

func (fileOrganizer *FileOrganizer) SetStrategy(strategy FileOperationStrategy) {
	fileOrganizer.Strategy = strategy
}

func (fileOrganizer *FileOrganizer) ExecuteStrategy() error {
	var fileManager *FileManager = NewFileManager(fileOrganizer.Logger)
	return fileOrganizer.Strategy.Execute(fileManager)
}
