package walker

import (
	"fmt"
	"os"
	"path/filepath"
)

func Walk(inputDir string) (filepathList []string) {

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		//fmt.Printf("%q\n", path)

		if filepath.Ext(path) == "avi" {
			filepathList = append(filepathList, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return filepathList

}
