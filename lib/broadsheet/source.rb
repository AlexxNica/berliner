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
    entries.reject! do |entry|
      entry.published < (Time.now - 3.days)
    end
    entries.map do |entry|
      open(entry.url).read
    end
  end

  def parse(html)
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
