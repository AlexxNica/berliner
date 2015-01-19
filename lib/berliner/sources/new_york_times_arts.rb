require "berliner/sources/new_york_times"

module Berliner
  class NewYorkTimesArts < NewYorkTimes
    feed "http://rss.nytimes.com/services/xml/rss/nyt/Arts.xml"
    title "The New York Times (Arts)"
    homepage "http://www.nytimes.com/arts"
  end
end
