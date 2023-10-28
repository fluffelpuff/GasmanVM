package fsysfile

type FileStat struct {
}

type FolderStat struct {
}

func (o *FileStat) GetAtributes() ([]struct{ Name, Value string }, error) {
	return nil, nil
}

func (o *FolderStat) GetAtributes() ([]struct{ Name, Value string }, error) {
	return nil, nil
}
