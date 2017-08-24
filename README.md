# go-octicons

[![Build Status](https://travis-ci.org/ZoltanLajosKis/go-octicons.svg?branch=master)](https://travis-ci.org/ZoltanLajosKis/go-octicons)
[![Go Report Card](https://goreportcard.com/badge/github.com/ZoltanLajosKis/go-octicons)](https://goreportcard.com/report/github.com/ZoltanLajosKis/go-octicons)
[![Coverage Status](https://coveralls.io/repos/github/ZoltanLajosKis/go-octicons/badge.svg?branch=master)](https://coveralls.io/github/ZoltanLajosKis/go-octicons?branch=master)

Go package for GitHub Octicons. Based on npm version 6.0.1.


## Install
The package can installed with `go get`. It contains the generated source
files, so no additional steps are necessary.
```sh
go get github.com/ZoltanLajosKis/go-octicons
```

## Usage
After importing the package icons can be accessed by symbol name or directly
as a variable. The list of symbols can also be retrieved.
```go
// Octicon by symbol.
octicons.Octicons("logo-github")

// Octicon directly as a variable.
octicons.LogoGithub

// List of symbols.
octicons.Symbols()
```

Octicons are represented by the following interface.
```go
type Octicon interface {
  // Symbol returns the symbol name, same as the key for that icon.
  Symbol() string
  // Keywords returns an array of keywords for the icon.
  Keywords() []string
  // Path returns the string representation of the path of the icon.
  Path() string
  // Options returns the attributes that will be added to the output tag.
  Options() Opts
  // Width returns the icon's true width, based on the svg view box width.
  Width() int
  // Height returns the icon's true height, based on the svg view box height.
  Height() int
  // ToSVG returns the string representation of the svg for the icon.
  ToSVG(Opts) string
  // ToSVGUse returns the string representation of the svg for the icon to use with sprites.
  ToSVGUse(Opts) string
}

type Opts map[string]string
```

Use the `ToSVG(Opts)` method to get a string of the `<svg>` tag for the icon.
```go
// With the default options.
icon.ToSVG(nil)

// Add more CSS classes to the <svg> tag.
icon.ToSVG(octicons.Opts{"class": "close"}),

// Add accessibility aria-label to the icon.
icon.ToSVG(octicons.Opts{"aria-label": "Close the window"}),

// Size the SVG icon larger using width & height independently or together.
icon.ToSVG(octiconsOpts{"width": "45"})
icon.ToSVG(octiconsOpts{"height": "60"})
icon.ToSVG(octiconsOpts{"width": "45", "height": "60"})
```

Use the `ToSVGUse(Opts)` method to get a string of the `<svg>` tag for the icon
with the `<use>` tag, for use with the spritesheet.
```go
// Also supports the same Options as ToSVG.
icon.ToSVGUse(nil)
```

The contents of static assets can also be retrieved as strings.
```go
// Minified Octicons CSS.
octicons.CSS
// Spritesheet containing all icons.
octicons.Spritesheet
```


## Testing
Testing compares the output of this package to the output of the
[octicons npm package][npm-octicons]. To run the tests, the following
executables must be available on the system:
- [npm][npm]: to retrieve the [octicons npm package][npm-octicons]
- [node][node]: to run the node.js octicons code
- [xmllint][xmllint]: to convert HTML output to canonical format for comparison

Tests can be executed using the following command.
```sh
make test
```


## Generate
The source code is generated from the [octicons npm package][npm-octicons]. To
generate, execute the following command.
```sh
make generate
```


## License
_SVG License:_ [SIL OFL 1.1](http://scripts.sil.org/OFL) (c) 2012-2016 GitHub, Inc.  
Applies to all SVG files

_Code License:_ [MIT](./LICENSE)  
Applies to all other files

When using the GitHub logos, be sure to follow the [GitHub logo guidelines](https://github.com/logos).

  
[node]: https://nodejs.org/
[npm-octicons]: https://www.npmjs.com/package/octicons
[npm]: https://www.npmjs.com/
[xmllint]: http://xmlsoft.org/xmllint.html
