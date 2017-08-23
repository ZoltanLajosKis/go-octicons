// +build ignore

package main

import (
	"log"

	as "github.com/ZoltanLajosKis/go-assets"
)

var (
	sources = []*as.Source{
		{"octicons",
			"https://registry.npmjs.org/packagename/-/octicons-6.0.1.tgz",
			&as.Checksum{as.MD5, "3ccfbbfb06fa12ccb0459d403a7da561"},
			&as.Archive{as.TarGz, as.ReMap(
				"(package/build/data.json|sprite.octicons.svg|octicons.min.css)",
				"${1}")}},
	}
)

func main() {
	fs, err := as.Retrieve(sources)
	check(err)

	items := getDataItems(fs)
	generate(fs, items)

	log.Println("octicons.go created.")
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
