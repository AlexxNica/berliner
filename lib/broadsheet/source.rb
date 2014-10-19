require "feedjira"
require "broadsheet/article"

class Source

  def self.fetch_articles
    entries = Feedjira::Feed.fetch_and_parse(@feed).entries
    entries.reject! do |entry|
      entry.published < (Time.now - (60*60*24))
    end
    articles = entries.map do |entry|
      Article.new(
          :title => entry.title,
          :author => entry.author,
          :content => entry.content,
          :published => entry.published,
          :url => entry.url,
          :source => @title,
          :style => @style
        )
    end
  end

  class << self
    attr_rw :title, :style, :feed
  end

end