class Article
  attr_accessor :title, :author, :content, :source

  def initialize(title, author, content, source)
    @title = title
    @author = author
    @content = content
    @source = source
  end
end
