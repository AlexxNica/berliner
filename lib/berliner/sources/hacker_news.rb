require "berliner/source"

module Berliner
  class HackerNews < Source
    feed "http://feeds.feedburner.com/newsyc100"
    title "Hacker News"
    homepage "http://news.ycombinator.com/"
  end
end
