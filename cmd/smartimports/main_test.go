package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func Test_processData_SeparateStdlib(t *testing.T) {
	src := `package main

import (
	"context"
	"github.com/pkg/errors"
	"os"
)

func main() {

}
`
	dst := `package main

import (
	"context"
	"os"

	"github.com/pkg/errors"
)

func main() {

}
`

	res, err := processData([]byte(src), getDefaultOpts())

	assert.NoError(t, err)
	assert.Equal(t, dst, string(res))
}

func Test_processData_NonPreformattedImports_ShouldBeProcessedCorrectly(t *testing.T) {
	src := `package main

	import (
	"context"
  "os"

		"github.com/pkg/errors"
		   
	"github.com/bradfitz/gomemcache"
	"fmt"
	)

func main() {

}
`
	dst := `package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bradfitz/gomemcache"
	"github.com/pkg/errors"
)

func main() {

}
`

	res, err := processData([]byte(src), getDefaultOpts())

	assert.NoError(t, err)
	assert.Equal(t, dst, string(res))
}
