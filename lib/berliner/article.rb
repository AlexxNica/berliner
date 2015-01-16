require "sanitize"

module Berliner
  # A Berliner article. The base object for every news feed item
  class Article
    attr_accessor :title,
      :body,
      :author,
      :date_published,
      :image,
      :location,
      :permalink,
      :source,
      :via,
      :other

    # Create a new {Article} object
    def initialize(
      title: title,
      body: body,
      author: nil,
      date_published: nil,
      image: nil,
      location: nil,
      permalink: nil,
      source: source,
      via: nil,
      other: {}
    )
      @title = sanitize(title)
      @body = sanitize(body, Sanitize::Config::BASIC)
      @author = sanitize(author)
      @date_published = sanitize(date_published)
      @image = sanitize(image)
      @location = sanitize(location)
      @permalink = sanitize(permalink)
      @source = sanitize(source)
      @via = sanitize(via)
      @other = other
    end

    private

    def sanitize(fragment, config = Sanitize::Config::RESTRICTED)
      return nil unless fragment
      Sanitize.fragment(fragment, config)
    end
  end
end
