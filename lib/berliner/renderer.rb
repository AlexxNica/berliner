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
      template = read_template(self.class.template)
      style = read_style(self.class.style)
      html = Erubis::Eruby.new(template).result({
        articles: articles,
        style: style
        })
      html_path = File.join(CONFIG_DIR, "berliner.html")
      File.write(html_path, html)
      begin
        system %{open "#{html_path}"}
      rescue
      end
    end

    # Read a CSS style file given its slug
    # @param [String] slug the style slug
    # @return [String] the contents of the style file
    def read_style(slug)
      filename = "#{slug.gsub(/-/, '_')}.css"
      begin
        template = File.read(File.join(Dir.home, ".berliner/styles", filename))
      rescue
        begin
          template = File.read(File.join(LIB_DIR, "berliner/styles", filename))
        rescue
          raise NameError,
            "The #{slug} CSS file was not found. " \
            "Make sure it is defined in styles/#{filename}"
        end
      end
      template
    end

    # Read an ERB template given its slug
    # @param [String] slug the template slug
    # @return [String] the contents of the template file
    def read_template(slug)
      filename = "#{slug.gsub(/-/, '_')}.erb"
      begin
        template = File.read(File.join(Dir.home, ".berliner/templates", filename))
      rescue
        begin
          template = File.read(File.join(LIB_DIR, "berliner/templates", filename))
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