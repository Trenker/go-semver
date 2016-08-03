# Go semver functions

This package contains types and functions to parse and check semver version numbers and constraints commonly used in package managers that rely on semver versions, like Ruby Gems or PHP Composer.

## Usage

There are only two access points in this package `NewVersion` and `NewConstraint`

### Version

`NewVersion` will return a Version type. The given string must be a valid semver version number. This means: 2 -  4 version digits and an optional suffix, which may be -alpha, -beta, -rc, -patch or -stable, optionally followed by a positive number.

A string that starts with `dev-` is considered a *rolling* version, as in a rolling release without explicit release versions. In this case, only equal comparisons can succeed.

### Constraint

`NewConstraint` takes a string which has one or more versions, optionally prepended by a comparator. The comparator itself can be prepended with an exclamation mark to negate the result of the comparison.

The following comparators are recognized:

Comparator | Explanation
-----------|------------
`=`        | Version must equal the given. Equals is set implicitly by omitting a comparator
`~`        | Version must be the same, but the last version digit can be higher. eg: `~1.2.3` matches `1.2.3` and `1.2.6`, but not `1.2.0` or `1.3`
`^`        | Version must be the same or higher, but not of a higher major version. eg: `^1.2.3` matches `1.2.6` and `1.6.0` but not `1.2.0` or `2.0`
`>`        | Version must be greater than given
`>=`       | Version must be equal or greater than given
`<`        | Version must be less than given
`<=`       | Version must be equal or less than given

A full constraint can contain multiple version constraints, connected by comma (`,`) or pipe (`|`). If the comma is used, the result and the result of the previous version must be true, when using the pipe, an *or* comparison is done.

eg:

* `>= 1.2 , < 2.0` matches `1.2.3` but not `1.0.0`
* `>= 1.2 | < 2.0` matches `1.2.3` and `1.0.0`

A constraint has two functions: `Match` and `MatchString`. The first will check if the version given matches the constraint. `MatchString` will do this with the version string, but will return *false*, not an error, if the version cannot be parsed.

### Example

```go
package main

import "github.com/garfieldius/go-semver"

v1 := semver.NewVersion("1.2.3")
v2 := semver.NewVersion("2.4.6")

c1 := semver.NewConstraint(">= 1.2.0, < 2.0")

c1.Match(v1) // true
c1.Match(v2) // false

c2 := semver.NewConstraint("^2.0.0 | 1.9.4-beta6")

c2.Match(v1) // false
c2.Match(v2) // true
```

## Credits

This code is highly inspired by the package [go-version](https://github.com/hashicorp/go-version), but is not a fork of it.

Many things in this package are based on the code in [composer/semver](https://github.com/composer/semver)

## License

Released under the [Apache License, Version 2.0](LICENSE)
