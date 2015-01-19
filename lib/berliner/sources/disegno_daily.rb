require "berliner/source"

module Berliner
  class DisegnoDaily < Source
    feed "http://feeds.feedburner.com/disegnofeed"
    title "Disegno Daily"
    homepage "http://www.disegnodaily.com/"

    def parse(entry)
      html = open(entry.url, allow_redirections: :safe).read
      doc = Nokogiri::HTML(html)
      title = doc.at_css("#main header hgroup h1").content
      image = doc.at_css("#main figure.media img")
              .attribute("src").content
      intro = doc.at_css("#main header .markdown").to_s
      main = doc.at_css("#main main .markdown").to_s
      body = "#{intro} #{main}"
      Article.new(
        title: title,
        author: "Disegno Daily",
        body: body,
        image: image,
        source: self.class.title,
        via: entry.via,
        permalink: entry.url
        )
    end
  end
end
