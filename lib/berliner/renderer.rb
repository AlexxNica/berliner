module Berliner
  # The base object for a Berliner renderer.  Each renderer should inherit from
  # {Renderer} and reimplement {Renderer#render} as necessary.
  # @abstract
  class Renderer

    # Create a new {Renderer} object
    def initialize(options = {})
      @options = options
    end

    # Render articles into a Berliner
    # @note Renderers usually output the Berliner to a file, but this
    #  behavior can be redefined in child classes
    # @param [Array<Article>] articles an array of {Article} objects
    # @return [void]
    def render(articles)
      articles.each do |article|
        puts "#{article.title} - #{article.author}"
      end

      return true # suppress printing articles array to console
    end
  end
end