require "berliner/sources/new_yorker"

module Berliner
  class NewYorkerCulture < NewYorker
    feed "http://www.newyorker.com/feed/culture"
    title "The New Yorker (Culture)"
    homepage "http://www.newyorker.com/culture"
  end
end