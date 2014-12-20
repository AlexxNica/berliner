require "berliner/sources/new_yorker"

module Berliner
  class NewYorkerTech < NewYorker
    feed "http://www.newyorker.com/feed/tech"
    title "The New Yorker (Tech)"
    homepage "http://www.newyorker.com/tech"
  end
end