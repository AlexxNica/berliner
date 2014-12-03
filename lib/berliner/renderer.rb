require "erubis"

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
      template = load(self.class.template)
      html = Erubis::Eruby.new(template).result({
        articles: articles
        })
      File.write("berliner.html", html)
    end

    def load(slug)
      filename = "#{slug.gsub(/-/, '_')}.erb"
      begin
        template = File.read("#{Dir.home}/.berliner/templates/#{filename}")
      rescue
        begin
          template = File.read("#{LIB_PATH}/berliner/templates/#{filename}")
        rescue
          raise NameError,
            "The #{slug} template was not found. " \
            "Make sure it is defined in templates/#{filename}"
        end
      end
      template
    end

    class << self
      # The ERB template to use to render articles
      # @note This attribute is set using a DSL
      # @example Define this attribute in child classes
      #   template "default"
      # @attribute [r]
      # @scope class
      # @return [String]
      attr_rw :template

      # The CSS style to apply to the Berliner html
      # @note This attribute is set using a DSL
      # @example Define this attribute in child classes
      #   style "default"
      # @attribute [r]
      # @scope class
      # @return [String]
      attr_rw :style
    end
  end
end