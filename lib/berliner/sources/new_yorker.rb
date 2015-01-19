require "berliner/source"
require "berliner/article"
require "nikkou"

module Berliner
  class NewYorker < Source
    feed "http://www.newyorker.com/feed/news"
    title "The New Yorker"
    homepage "http://www.newyorker.com/"

    def parse(entry)
      html = open(entry.url, allow_redirections: :safe).read
      doc = Nokogiri::HTML(html)
      title = doc.at_css("hgroup h1").content || ""
      author = doc.attr_equals("itemprop", "name author").first.content || ""
      image = doc.at_css("#articleBody figure img").attribute("data-src").content || nil
      body_node = doc.at_css("#articleBody")
      body_node.css("figure").remove
      body_node.css(".content-ad-wrapper").remove
      body_node.css(".social-buttons").remove
      body = body_node.to_s || ""
      Article.new(
        title: title,
        author: author,
        body: body,
        image: image,
        source: self.class.title,
        via: entry.via,
        permalink: entry.url
        )
    end
  end
end
