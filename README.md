# Smart imports

If you want to format imports in your golang source files according to the standards of your team, this tool might be handy.

Standard tool `goimports` doesn't regroup imports. For example if you have a file
```
package main

import (
    "context"
    "os"

    "github.com/pkg/errors"

    "github.com/bradfitz/gomemcache"
)
```
`goimports` wouldn't do anything with it. But `smartimports` would group imports into 2 groups:

```
package main

import (
    "context"
    "os"

    "github.com/pkg/errors"
    "github.com/bradfitz/gomemcache"
)
```

# Installation

For go 1.16: `go install github.com/pav5000/smartimports/cmd/smartimports@latest`

For earlier versions of go: `go get github.com/pav5000/smartimports/cmd/smartimports`

# Usage

Simple example:

`smartimports` (formats files in the current dir and all subdirs)

Complex example:
`smartimports -path ./my-project -exclude ./my-project/pb_pkg -local github.com/someuser/my-project`

With option `-path` you can specify a directory which should be processed (by default smartimports process current directory).

With option `-exclude` you can specify prefixes for paths which shouldn't be processed. Prefixes should match the style of `-path` option. You can check details with `-v`.
If your `-path` is absolute, `-exclude` should be absolute too.

With option `-local` you can specify a comma-separated list of prefixes for packages which should be grouped into 3rd group. Some teams put their local packages into it.

With option `-v` you can see verbose output of traversing through directories and files.
