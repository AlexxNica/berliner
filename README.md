# Broadsheet

Broadsheet is a Ruby gem and CLI that compiles a daily digest of online news in a beautiful format.

# Installation

```
$ gem install broadsheet
```

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

# Adding a Source

All sources inherit from the base Broadsheet `Source` object.  Here are the methods and attributes a custom source can implement:

```ruby
# lib/broadsheet/sources/my-source.rb
require "broadsheet/source"

class MySource < Source
  # The feed string is passed to Feedjira to download entries
  feed "http://feeds.feedburner.com/<my-source>"
  title "My Source"
  # Specify a custom stylesheet for formatting
  style "source"

  # define `self.fetch` to specify how the feed string returns
  # a list of html pages.  By default each feed entry less than
  # a day old is downloaded.

  # the `self.parse` method takes html and returns an Article object
  def self.parse(html)
    require "nokogiri"
    require "broadsheet/article"

    doc = Nokogiri::HTML(html)

    Article.new(
      title: doc.css('title'),
      author: doc.at("meta[name='author']")['content'],
      content: doc.css('article'),
      published: doc.css('date-published'),
      url: doc.css('permalink'),
      source: @title,
      style: @style
      )
  end
end
```

# Supported Ruby Versions

Broadsheet requires 2.1.2 or higher.
