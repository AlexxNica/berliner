require "berliner/source"

module Berliner
  class Vox < Source
    feed "http://www.vox.com/rss/index.xml"
    title "Vox"
    homepage "http://www.vox.com/"

    # Vox has full content in the rss feed so we don't have to scrape
    def articles
      @entries = fetch
      articles = @entries.map do |entry|
        Article.new(
          title: entry.title,
          author: entry.author,
          body: entry.content,
          source: self.class.title
          )
      end
    end
  end
end
