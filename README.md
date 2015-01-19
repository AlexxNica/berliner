# Berliner

> WARNING: Berliner is experimental software under active development.  There are few tests. The CLI commands, architecture, APIs, and source code layout may change at any moment.

> **Berliner will be fit for public use at v0.1.0** (follow progress [here](https://github.com/s3ththompson/berliner/issues/10) and [here](https://github.com/s3ththompson/berliner/issues/11).)

[![Gitter](https://img.shields.io/badge/GITTER-JOIN%20CHAT%20%E2%86%92-brightgreen.svg?style=flat)](https://gitter.im/s3ththompson/berliner)
[![Gem Version](https://badge.fury.io/rb/berliner.svg)](http://badge.fury.io/rb/berliner)
[![Build Status](https://travis-ci.org/s3ththompson/berliner.svg?branch=master)](https://travis-ci.org/s3ththompson/berliner)
[![Coverage Status](https://img.shields.io/coveralls/s3ththompson/berliner.svg?style=flat)](https://coveralls.io/r/s3ththompson/berliner?branch=master)
[![Inline docs](http://inch-ci.org/github/s3ththompson/berliner.svg?branch=master)](http://inch-ci.org/github/s3ththompson/berliner)
[![Code Climate](https://codeclimate.com/github/s3ththompson/berliner/badges/gpa.svg)](https://codeclimate.com/github/s3ththompson/berliner)

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
$ berliner add new-yorker-books
$ berliner add five-thirty-eight

# list your added sources
$ berliner list

# Read your daily Berliner
$ berliner read
```

# Advanced Usage

Build a profile of your favorite news sources and save it to `~/.berliner/profile.yml`.  For a list of supported sources, see [here](https://github.com/s3ththompson/berliner/tree/master/lib/berliner/sources).  Choose a renderer from the supported renderers [here](https://github.com/s3ththompson/berliner/tree/master/lib/berliner/renderers).  Add any number of filters from the supported filters [here](https://github.com/s3ththompson/berliner/tree/master/lib/berliner/filters).

``` yaml
#  Example ~/.berliner/profile.yml
sources:
  - designo-daily
  - new-york-times
  - new-yorker
filters:
  - per-source-limit
renderer: utilitarian
```

Build a Berliner from the config settings in your profile.

```sh
$ berliner read
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
  homepage "http://www.my-source.com/"

  # define `fetch` to specify how the feed string returns
  # a list of html pages.

  # the `parse` method takes a Berliner::Feed::FeedEntry object and returns an Article object
  def parse(entry)
    require "nokogiri"
    require "berliner/article"
    require "open-uri"

    html = open(entry.url).read
    doc = Nokogiri::HTML(html)

    Article.new(
      title: doc.at_css("title").content,
      author: doc.at("meta[name='author']")["content"],
      content: doc.at_css("article").to_s,
      published: doc.at_css(".date-published").content,
      url: doc.at_css("#permalink").content,
      source: self.class.title,
      via: entry.via
      )
  end
end
```

# API Documentation

[Berliner](http://www.rubydoc.info/gems/berliner/) on RubyDoc.info

# Supported Ruby Versions

Berliner requires Ruby 2.0.0 or greater.  It is tested against Ruby 2.0.0, 2.1.1, and 2.1.2.
