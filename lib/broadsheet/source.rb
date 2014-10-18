require 'feedjira'

class Source

  def self.download
    new(@feed, @title, @style).download
  end

  def initialize(feed, title, style)
    @feed = feed
    @title = title
    @style = style
  end

  def download
    f = Feedjira::Feed.fetch_and_parse(@feed)
    f.entries
  end

  class << self
    attr_rw :title, :style

    def feed val
      @feed = val
    end
  end

end