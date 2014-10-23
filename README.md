# Berliner

[![Gem Version](https://badge.fury.io/rb/berliner.svg)](http://badge.fury.io/rb/berliner)
[![Build Status](https://travis-ci.org/s3ththompson/berliner.svg?branch=master)](https://travis-ci.org/s3ththompson/berliner)
[![Coverage Status](https://img.shields.io/coveralls/s3ththompson/berliner.svg)](https://coveralls.io/r/s3ththompson/berliner?branch=master)

<img align="right" height="250" src="http://i.imgur.com/4BEEus9.png" alt="Berliner">

> Daily digests of online news

Berliner is a Ruby gem and UNIX-style CLI that compiles a daily digest of online news in a beautiful format. More than just another RSS reader, Berliner offers an elegant solution for pruning infinite streams of articles down to just the best content. Berliner scrapes each news item directly, and sends article compilations to a variety of renderers (PDF, HTML, Markdown, etc.). Don't like how your digest looks? Missing your favorite news source? Berliner is built to be hackable, extensible, and customizable.

# Installation

```sh
$ gem install berliner
```

# CLI Usage

```sh
# list all sources
$ berliner search

# add your preferred sources
$ berliner add new-york-times
$ berliner add five-thirty-eight

# list your added sources
$ berliner list

# Read your daily Berliner
$ berliner read
# Or specify a renderer
$ berliner read --console

```

# Module Usage

Build a profile of your favorite news sources and save it to `~/.berliner/profile.yml`.  For a list of supported sources, see [here](https://github.com/s3ththompson/berliner/tree/master/lib/berliner/sources).  Choose a renderer from the supported renderers [here](https://github.com/s3ththompson/berliner/tree/master/lib/berliner/renderers).

``` yaml
#  Example ~/.berliner/profile.yml
sources:
  - designo-daily
  - new-york-times
  - new-yorker
renderer: pdf
```

Build a Berliner from the sources in your profile.

``` ruby
require "berliner"
Berliner.read()
```

# Adding a Source

All sources inherit from the base Berliner `Source` object.  Here are the methods and attributes a custom source can implement:

```ruby
# lib/berliner/sources/my_source.rb
require "berliner/source"

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
    require "berliner/article"

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

# API Documentation

[Berliner](http://www.rubydoc.info/gems/berliner/) on RubyDoc.info

# Supported Ruby Versions

Berliner requires Ruby 2.0.0 or great.  It is tested against Ruby 2.0.0, 2.1.1, and 2.1.2.
