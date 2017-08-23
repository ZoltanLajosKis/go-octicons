// +build ignore

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type dataItems map[string]*dataItem

type dataItem struct {
	keywords []string
	path     string
	height   int
	width    int
}

func getDataItems(fs http.FileSystem) dataItems {
	f, err := fs.Open("package/build/data.json")
	check(err)

	d, err := ioutil.ReadAll(f)
	check(err)

	items := make(dataItems)
	err = json.Unmarshal(d, &items)
	check(err)

	return items
}

func (items dataItems) UnmarshalJSON(data []byte) error {
	aux := make(map[string]*dataItem)
	err := json.Unmarshal(data, &aux)

	if err == nil {
		for k, v := range aux {
			items[k] = v
		}
	}

	return err
}

func (item *dataItem) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Keywords []string `json:"keywords"`
		Path     string   `json:"path"`
		Height   int      `json:"height,string"`
		Width    int      `json:"width,string"`
	}{}

	err := json.Unmarshal(data, &aux)

	if err == nil {
		item.keywords = aux.Keywords
		item.path = aux.Path
		item.height = aux.Height
		item.width = aux.Width
	}

	return err
}
