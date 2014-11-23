require "berliner/source"

module Berliner
  class Slate < Source
    feed "http://feeds.slate.com/slate"
    title "Slate"
    style "slate"
  end
end