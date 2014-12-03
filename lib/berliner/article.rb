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
      title: nil,
      body: nil,
      author: nil,
      date_published: nil,
      image: nil,
      location: nil,
      permalink: nil,
      source: nil,
      other: {}
      )
      raise "title required" if title.nil?
      raise "body required" if body.nil?
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