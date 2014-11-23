require "berliner/source"

module Berliner
  class FiveThirtyEight < Source
    feed "http://fivethirtyeight.com/features/feed/"
    title "FiveThirtyEight"
    style "fivethirtyeight"
  end
end