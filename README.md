# Broadsheet

Broadsheet is a Ruby gem and CLI that compiles a daily digest of online news in a beautiful format.

# Installation

  $ gem install broadsheet

# Usage

Build a profile of your favorite news sources and save it to `~/.broadsheet/profile.yaml`.  For a list of supported sources, see [here](https://github.com/s3ththompson/broadsheet/tree/master/lib/broadsheet/sources).

``` yaml
#  Example ~/.broadsheet/profile.yaml
sources:
  - designo-daily
  - new-york-times
  - new-yorker
```

Build a Broadsheet from the sources in your profile.

``` ruby
require "broadsheet"
Broadsheet.build()
```

# Supported Ruby Versions

Broadsheet requires 2.1.2 or higher.
