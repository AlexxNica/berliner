require "berliner/source"

module Berliner
  class NewYorker < Source
    feed "http://www.newyorker.com/feed/news"
    title "The New Yorker"
    homepage "http://www.newyorker.com/"
  end
end