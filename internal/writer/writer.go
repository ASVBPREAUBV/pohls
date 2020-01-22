package writer

import (
	"fmt"
	"github.com/ASVBPREAUBV/pohls/internal/resolver"
	"io/ioutil"
	"path"
	"strings"
)

func writer(targetDir string, mediaList []resolver.Media) {

	for _, media := range mediaList {
		sourceDir := path.Join(targetDir, string(media.MediaType))
		if media.MediaType == resolver.MediaTypeSeries {
			sourceDir = path.Join(sourceDir, fmt.Sprintf("%s", strings.Replace(media.Title, " ", "_", -1)), fmt.Sprintf("season_%02d", media.Season), fmt.Sprintf("s%02d", media.Season), fmt.Sprintf("%s s%02de%02d", strings.Replace(media.Title, " ", "_", -1), media.Season, media.Episode))
		} else {
			sourceDir = path.Join(sourceDir, fmt.Sprintf("%s %d", strings.Replace(media.Title, " ", "_", -1), media.Year))
		}
		writeFile(media.SourceFilePath, sourceDir)
	}

}

func writeFile(source, target string) {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(target, input, 0644)
	if err != nil {
		fmt.Println("Error creating", target)
		return
	}
}
