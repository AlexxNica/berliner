require "feedjira"
require "ruby-readability"
require "open-uri"
require "broadsheet/article"

class Source

  def self.fetch
    entries = Feedjira::Feed.fetch_and_parse(@feed).entries
    entries.reject! do |entry|
      entry.published < (Time.now - (60*60*24))
    end
    entries.map do |entry|
      open(entry.url).read
    end
  end

  def self.parse(html)
    document = Readability::Document.new(html)
    Article.new(
      :title => document.title,
      :author => document.author,
      :content => document.content,
      :source => @title,
      :style => @style
      )
  end

  def self.articles
    @entries = self.fetch
    articles = @entries.map do |entry|
      self.parse(entry)
    end
  end

  class << self
    attr_rw :title, :style, :feed
  end

end