require "berliner/source"

module Berliner
  class Vox < Source
    feed "http://www.vox.com/rss/index.xml"
    title "Vox"
    style "vox"

    # Vox has full content in the rss feed so we don't have to scrape
    def articles
      @entries = fetch
      articles = @entries.map do |entry|
        Article.new(
          title: entry.title,
          author: entry.author,
          content: entry.content,
          source: "Vox",
          style: "vox"
          )
      end
    end
  end
end