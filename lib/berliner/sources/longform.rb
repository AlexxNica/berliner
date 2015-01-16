require "berliner/source"
require "nikkou"

module Berliner
  class Longform < Source
    feed "http://longform.org/feed.rss"
    title "Longform"
    homepage "http://longform.org/"

    def fetch
      feedjira_entries = Feedjira::Feed.fetch_and_parse(self.class.feed).entries
      feedjira_entries.map do |e|
        url = find_url(e.content)
        url ? Feed::FeedEntry.new(url, self.class.title) : nil
      end.compact
    end

    private

    def find_url(content)
      doc = Nokogiri::HTML(content)
      link = doc.search("a").text_includes("Full Story").first
      return link ? link["href"] : false
    end
  end
end