require "berliner/source"

module Berliner
  class DisegnoDaily < Source
    feed "http://feeds.feedburner.com/disegnofeed"
    title "Disegno Daily"
  end
end