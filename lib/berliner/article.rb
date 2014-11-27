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
      other: {}
    )
      @title = title
      @body = body
      @author = author
      @date_published = date_published
      @image = image
      @location = location
      @permalink = permalink
      @source = source
      @other = other
    end
  end
end
