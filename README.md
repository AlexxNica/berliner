# Broadsheet
[![Build Status](https://travis-ci.org/s3ththompson/broadsheet.svg?branch=master)](https://travis-ci.org/s3ththompson/broadsheet)
[![Coverage Status](https://img.shields.io/coveralls/s3ththompson/broadsheet.svg)](https://coveralls.io/r/s3ththompson/broadsheet?branch=master)

Broadsheet is a Ruby gem and CLI that compiles a daily digest of online news in a beautiful format.

# Installation

```sh
$ gem install broadsheet
```

# CLI Usage

```sh
# list all sources
$ broadsheet search

# add your preferred sources
$ broadsheet add new-york-times
$ broadsheet add five-thirty-eight

# list your added sources
$ broadsheet list

# Read your daily Broadsheet
$ broadsheet read
# Or specify a renderer
$ broadsheet read --console

```

# Module Usage

Build a profile of your favorite news sources and save it to `~/.broadsheet/profile.yml`.  For a list of supported sources, see [here](https://github.com/s3ththompson/broadsheet/tree/master/lib/broadsheet/sources).  Choose a renderer from the supported renderers [here](https://github.com/s3ththompson/broadsheet/tree/master/lib/broadsheet/renderers).

``` yaml
#  Example ~/.broadsheet/profile.yml
sources:
  - designo-daily
  - new-york-times
  - new-yorker
renderer: pdf
```

Build a Broadsheet from the sources in your profile.

``` ruby
require "broadsheet"
Broadsheet.read()
```

# Adding a Source

All sources inherit from the base Broadsheet `Source` object.  Here are the methods and attributes a custom source can implement:

```ruby
# lib/broadsheet/sources/my_source.rb
require "broadsheet/source"

class MySource < Source
  # The feed string is passed to Feedjira to download entries
  feed "http://feeds.feedburner.com/<my-source>"
  title "My Source"
  # Specify a custom stylesheet for formatting
  style "source"

  # define `fetch` to specify how the feed string returns
  # a list of html pages.  By default each feed entry less than
  # a day old is downloaded.

  # the `parse` method takes html and returns an Article object
  def parse(html)
    require "nokogiri"
    require "broadsheet/article"

    doc = Nokogiri::HTML(html)

    Article.new(
      title: doc.css("title"),
      author: doc.at("meta[name='author']")["content"],
      content: doc.css("article"),
      published: doc.css("date-published"),
      url: doc.css("permalink"),
      source: self.class.title,
      style: self.class.style
      )
  end
end
```

# Supported Ruby Versions

Broadsheet is tested on Ruby 2.1.0 and 2.1.2.
