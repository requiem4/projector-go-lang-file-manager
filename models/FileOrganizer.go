package models

import "log"

type FileOrganizer struct {
	Strategy FileOperationStrategy
	logger   *log.Logger
}

func (fileOrganizer *FileOrganizer) SetStrategy(strategy FileOperationStrategy) {
	fileOrganizer.Strategy = strategy
}

func (fileOrganizer *FileOrganizer) ExecuteStrategy() error {
	var fileManager *FileManager = NewFileManager(fileOrganizer.logger)
	return fileOrganizer.Strategy.Execute(fileManager)
}
