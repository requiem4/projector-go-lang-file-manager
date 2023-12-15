package models

type FileOrganizer struct {
	Strategy FileOperationStrategy
}

func (fo *FileOrganizer) SetStrategy(strategy FileOperationStrategy) {
	fo.Strategy = strategy
}

func (fo *FileOrganizer) ExecuteStrategy() error {
	var fileManager *FileManager = &FileManager{}
	return fo.Strategy.Execute(fileManager)
}
