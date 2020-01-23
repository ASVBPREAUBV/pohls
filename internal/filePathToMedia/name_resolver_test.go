package filePathToMedia_test

import (
	"encoding/json"
	"fmt"
	"github.com/ASVBPREAUBV/pohls/internal/resolver"
	"io/ioutil"
	"testing"
)

func TestNameToTmdb(t *testing.T) {
	dat, err := ioutil.ReadFile("../../tests/filenames.json")
	if err != nil {
		t.Error(err)
	}

	var filenames []string
	err = json.Unmarshal(dat, &filenames)
	if err != nil {
		t.Error(err)
	}

	media := filePathToMedia.FilenameCleaner(filenames)

	for _, m := range media {
		fmt.Println(m.Title)
	}
}
