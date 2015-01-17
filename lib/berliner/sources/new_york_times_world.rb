require "berliner/sources/new_york_times"

module Berliner
  class NewYorkTimesWorld < NewYorkTimes
    feed "http://rss.nytimes.com/services/xml/rss/nyt/World.xml"
    title "The New York Times (World)"
    homepage "http://www.nytimes.com/world"
  end
end