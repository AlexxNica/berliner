require "feedjira"
require "ruby-readability"
require "open-uri"
require "broadsheet/article"
require "active_support/core_ext"

class Source

  def initialize(options = {})
    @options = options
  end

  def fetch
    entries = Feedjira::Feed.fetch_and_parse(self.class.feed).entries

    # For now, take articles published in the last 3 days, up to a max of 3 articles
    entries.reject{ |entry| entry.published < (Time.now - 3.days) }.take(3)
  end

  def parse(entry)
    html = open(entry.url).read
    document = Readability::Document.new(html)
    Article.new(
      title: document.title,
      author: document.author,
      content: document.content,
      source: self.class.title,
      style: self.class.style
      )
  end

  def articles
    @entries = fetch
    articles = @entries.map do |entry|
      parse(entry)
    end
  end

  class << self
    attr_rw :title, :style, :feed
  end

end
