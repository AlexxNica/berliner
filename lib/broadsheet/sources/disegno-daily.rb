require 'broadsheet/source'

class DisegnoDaily < Source
  feed 'http://feeds.feedburner.com/disegnofeed'
  title 'Disegno Daily'
  style 'disegno'
end