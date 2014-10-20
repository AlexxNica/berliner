require "prawn"

class Renderer

  def initialize(options = {})
    @options = options
  end

  # The render function takes an array of Article objects and an options hash,
  # which can be in an arbitrary format depending on the subclass.
  #
  # This is an example renderer which just prints the articles to the console.
  def render(articles)
    articles.each do |article|
      puts "#{article.title} - #{article.author}"
    end

    return true # suppress printing articles array to console
  end
end
