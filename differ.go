// Copyright © 2022 Mark Summerfield. All rights reserved.
// License: Apache-2.0

package differ

import (
    "fmt"
    _ "embed"
    )

//go:embed Version.dat
var Version string

func Hello() string {
    return fmt.Sprintf("Hello differ v%s", Version)
}
