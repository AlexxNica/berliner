package renderers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/s3ththompson/berliner/content"
)

func ToJSON(filename string) (string, func([]content.Post)) {
	return "JSON", func(posts []content.Post) {
		p, err := json.Marshal(posts)
		if err != nil {
			return
		}
		_ = ioutil.WriteFile(filename, p, 0644)
	}
}
