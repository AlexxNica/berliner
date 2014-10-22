class Article
  attr_accessor :title, :author, :content, :published, :url, :source, :style, :custom

  def initialize(args={})
    args.each do |attr, val|
      instance_variable_set("@#{attr}", val)
    end if args
  end
end