package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/imports"
)

func getDefaultOpts() *imports.Options {
	return &imports.Options{
		TabIndent:  true,
		FormatOnly: true,
	}
}

func Test_processData_SeveralImportBlocksToOne(t *testing.T) {
	src := `package main

import "context"
import "os"

func main() {

}
`
	dst := `package main

import (
	"context"
	"os"
)

func main() {

}
`

	res, err := processData([]byte(src), getDefaultOpts())

	assert.NoError(t, err)
	assert.Equal(t, dst, string(res))
}

func Test_processData_MergeImportSections(t *testing.T) {
	src := `package main

import (
	"context"

	"os"
)

func main() {

}
`
	dst := `package main

import (
	"context"
	"os"
)

func main() {

}
`

	res, err := processData([]byte(src), getDefaultOpts())

	assert.NoError(t, err)
	assert.Equal(t, dst, string(res))
}
