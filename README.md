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
`goimports` wouldn't fo anything with it. But `smartimports` would group imports into 2 groups:

```
package main

import (
    "context"
    "os"

    "github.com/pkg/errors"
    "github.com/bradfitz/gomemcache"
)
```

With option `-local` you can specify a comma-separated list of prefixes for packages which should be grouped into 3rd group. Some teams put their local packages into it.
