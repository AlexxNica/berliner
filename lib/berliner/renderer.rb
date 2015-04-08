require "erubis"
require "fileutils"
require "uri"
require "pathname"
require "parallel"
require "timeout"
require "berliner/loader"

module Berliner
  # The base object for a Berliner renderer.  Each renderer should inherit from
  # {Renderer} and reimplement {Renderer#render} as necessary.
  # @abstract
  class Renderer
    # Path to the output berliner.html file
    BERLINER_HTML = File.join(CONFIG_DIR, "berliner.html")
    # Path to the assets output folder berliner_files
    BERLINER_HTML_FILES = File.join(CONFIG_DIR, "berliner_files")

    # Create a new {Renderer} object
    def initialize(options = {})
      @options = options
      clean_up
      make_files_dir
    end

    # Render articles into a Berliner
    # @note Renderers usually output the Berliner to a file, but this
    #  behavior can be redefined in child classes
    # @param [Array<Article>] articles an array of {Article} objects
    # @return [void]
    def render(articles)
      template = read_template(self.class.template)
      style = read_style(self.class.style)
      articles = save_images(articles)
      html = Erubis::Eruby.new(template).result(
        articles: articles,
        style: style
        )
      File.write(BERLINER_HTML, html)
      BERLINER_HTML
    end

    # Read a CSS style file given its slug
    # @param [String] slug the style slug
    # @return [String] the contents of the style file
    def read_style(slug)
      filename = "#{slug.deslugify}.css"
      begin
        style = Loader.read_file(File.join("berliner",
                                           "assets", "styles", filename))
      rescue
        raise NameError,
              "The #{slug} CSS file was not found. " \
              "Make sure it is defined in assets/styles/#{filename}"
      end
      style
    end

    # Read an ERB template given its slug
    # @param [String] slug the template slug
    # @return [String] the contents of the template file
    def read_template(slug)
      filename = "#{slug.deslugify}.erb"
      begin
        template = Loader.read_file(File.join("berliner",
                                              "assets", "templates", filename))
      rescue
        raise NameError,
              "The #{slug} template was not found. " \
              "Make sure it is defined in assets/templates/#{filename}"
      end
      template
    end

    # Save all images in articles to disk and replace URL with relative filename
    # @param [Array<Article>] articles an array of {Article} objects
    # @return [Array<Article>] an array of {Article} objects with image
    #   attribute altered
    def save_images(articles)
      Parallel.map(articles, in_threads: 10) do |article|
        article.image = save_image(article.image)
        article
      end
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

    private

    # Clean up old berliner.html and berliner_files folder
    # @return [void]
    def clean_up
      FileUtils.remove_entry_secure(
               BERLINER_HTML_FILES) if File.exist?(BERLINER_HTML_FILES)
      FileUtils.remove_entry_secure(BERLINER_HTML) if File.exist?(BERLINER_HTML)
    end

    # Create a 'berliner_files' directory in the config dir
    # @return [void]
    def make_files_dir
      FileUtils.mkdir_p(BERLINER_HTML_FILES) unless File.directory?(
                                                         BERLINER_HTML_FILES)
    end

    # Download an image given its url, save to disk, and return relative
    #   file location
    # @param [String, nil] url the image url
    # @return [String, nil] relative file location of image on disk,
    #   or nil if no image
    def save_image(url)
      return nil unless url
      uri = URI.parse(url)
      basename = File.basename(URI.unescape(uri.path))
      file = File.join(BERLINER_HTML_FILES, basename)
      begin
        # Timeout image download after 10 seconds
        Timeout.timeout(10) do
          File.open(file, "wb") { |f| f.write(open(uri).read) }
        end
        relative = Pathname.new(file)
                   .relative_path_from(Pathname.new(CONFIG_DIR)).to_s
        return relative
      rescue
        return nil
      end
    end
  end
end
