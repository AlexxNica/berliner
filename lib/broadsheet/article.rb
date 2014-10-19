require 'ruby-readability'
require 'open-uri'

class Article
  attr_accessor :title, :author, :content, :published, :url, :source, :style

  def self.filter(articles, profile)
    articles.shuffle.first(5)
  end

  def initialize(args={})
    args.each do |attr, val|
      instance_variable_set("@#{attr}", val)
    end if args

    @html = open(@url).read
    @document = Readability::Document.new(@html)
    @content = @document.content
    @author = @document.author
    @title = @document.title
  end
end
