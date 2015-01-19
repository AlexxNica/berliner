require "berliner/sources/new_yorker"

module Berliner
  class NewYorkerBooks < NewYorker
    feed "http://www.newyorker.com/feed/books"
    title "The New Yorker (Books)"
    homepage "http://www.newyorker.com/books"
  end
end
