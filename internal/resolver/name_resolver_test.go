package resolver_test

import (
	"encoding/json"
	"github.com/ASVBPREAUBV/pohls/internal/resolver"
	"io/ioutil"
	"testing"
)

func TestNameToTmdb(t *testing.T) {
	dat, err := ioutil.ReadFile("/Users/paul/go/src/github.com/ASVBPREAUBV/pohls/tests/filenames.json")
	if err != nil {
		t.Error(err)
	}

	var filenames []string
	err = json.Unmarshal(dat, &filenames)
	if err != nil {
		t.Error(err)
	}

	//fmt.Print(filenames)

	for _, filename := range filenames {
		resolver.FilePathToTmdbCollection(filename)
	}
}
