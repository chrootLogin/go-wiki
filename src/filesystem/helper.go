package filesystem

import (
	"os"
	"log"
	"io/ioutil"
)

func readFile(fs *filesystem, path string) (*File, error) {
	// Get FileInfo
	fileinfo, err := fs.fs.Stat(path)
	if err != nil {
		return nil, err
	}

	if fileinfo.IsDir() {
		return nil, ErrIsDir
	}

	// Open file
	file, err := fs.fs.OpenFile(path, os.O_RDONLY, fs.perms)
	if err != nil {
		log.Println("open file: " + err.Error())
		return nil, err
	}

	// Read file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("read file: " + err.Error())
		return nil, err
	}

	// Close file
	err = file.Close()
	if err != nil {
		log.Println("close file: " + err.Error())
		return nil, err
	}

	return &File{
		Content: string(data),
	}, nil
}