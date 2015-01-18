require "berliner/source"
require "berliner/article"
require "mechanize"
require "nikkou"

module Berliner
  class NewYorkTimes < Source
    feed "http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml"
    title "The New York Times"
    homepage "http://www.nytimes.com/"

    def auth
      return false unless (creds["username"] && creds["password"])
      @mech = Mechanize.new
      page = @mech.get('https://myaccount.nytimes.com/auth/login')
      form = page.forms.first
      form.set_fields(
        userid: creds["username"],
        password: creds["password"]
        )
      form.click_button
      return true
    end

    def parse(entry)
      if authenticated
        page = @mech.get(entry.url)
        doc = page.parser
      else
        # TODO: Recognize the NYT paywall instead of assuming we hit it.
        return nil
      end
      title = doc.at("meta[name='hdl']")["content"] || ""
      author = doc.at("meta[name='author']")["content"] || ""
      begin
        image = doc.at_css(".lede-container figure .image img").attribute("data-mediaviewer-src").content || nil
      rescue
        image = nil
      end
      body_node = doc.at_css("#story-body")
      body = body_node.css("p.story-body-text.story-content").map do |p|
        p.to_s
      end.join("") || ""
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