# A Berliner article. The base object for every news feed item
class Article
  attr_accessor :title, :author, :content, :published, :url, :source, :style, :custom

  # Create a new {Article} object
  def initialize(options = {})
    options.each do |attr, val|
      instance_variable_set("@#{attr}", val)
    end if options
  end
end