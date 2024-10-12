# versioncompare - Go Library for Version Comparison

The `versioncompare` library provides functions for comparing version strings in Go.

## Installation

To install the library, use the `go install` command:

go install  github.com/ansxuman/versioncompare


## Usage

1. Import the `versioncompare` package in your Go code:

```go
import "github.com/ansxuman/versioncompare"

To compare two versions, use the IsGreater, IsLess, or IsEqual functions from the versioncompare package:

currentVersion := "1.2.3"
requiredVersion := "1.2.0"

isGreater, err := versioncompare.IsGreater(currentVersion, requiredVersion)
if err != nil {
    // Handle the error
}

if isGreater {
    // Do something if currentVersion is greater than requiredVersion
}

isLess, err := versioncompare.IsLess(currentVersion, requiredVersion)
if err != nil {
    // Handle the error
}

if isLess {
    // Do something if currentVersion is less than requiredVersion
}

isEqual, err := versioncompare.IsEqual(currentVersion, requiredVersion)
if err != nil {
    // Handle the error
}

if isEqual {
    // Do something if currentVersion is equal to requiredVersion
}
```

# License
This library is released under the MIT License. See LICENSE for more information.

