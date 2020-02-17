package files

import (
	"fmt"
	"os"
	"path/filepath"
)

var allowedFileExtentions = []string{"avi", "mp4", "mkv", "mpg"}

func Walk(inputDir string) (filepathList []string) {

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		//fmt.Println(path, filepath.Ext(path))

		for _, ex := range allowedFileExtentions {
			if filepath.Ext(path) == "."+ex {
				filepathList = append(filepathList, filepath.Base(path))
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return filepathList

}
