require "berliner/sources/new_york_times"

module Berliner
  class NewYorkTimesTech < NewYorkTimes
    feed "http://rss.nytimes.com/services/xml/rss/nyt/Technology.xml"
    title "The New York Times (Technology)"
    homepage "http://www.nytimes.com/technology"
  end
end
