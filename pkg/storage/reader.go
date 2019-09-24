package storage

import (
	"fmt"
	"io/ioutil"
)

// Storage struct represents the local storage state. This local state
// represents what the MongoDB database should eventually become. This
// object contains all of the local files from one database, including
// the collections.
type Storage struct {
	// Documents is a double map, which represents the local storage data.
	// The first key in the map is the collection key, which is loaded by
	// the folder name. The second key is the filename (the document ID)
	// which is loaded by the filename. The slice of bytes represents the
	// content of that file.
	Documents map[string]map[string][]byte
}

// ReadStorage will read the specified directory, and serialize the found
// directories and files into a Storage object. If anything fails, error
// is returned.
func ReadStorage(storDir string) (*Storage, error) {
	dirs, err := ioutil.ReadDir(storDir)
	if err != nil {
		return nil, err
	}
	stor := &Storage{
		Documents: make(map[string]map[string][]byte, len(dirs)),
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", storDir, dir.Name()))
		if err != nil {
			return nil, err
		}
		stor.Documents[dir.Name()] = make(map[string][]byte, len(files))
		for _, file := range files {
			cont, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", storDir, dir.Name(), file.Name()))
			if err != nil {
				return nil, err
			}
			stor.Documents[dir.Name()][file.Name()] = cont
		}
	}

	return stor, nil
}
