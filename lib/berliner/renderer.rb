require "erubis"
require "fileutils"
require "uri"
require "pathname"
require "parallel"
require "timeout"

module Berliner
  # The base object for a Berliner renderer.  Each renderer should inherit from
  # {Renderer} and reimplement {Renderer#render} as necessary.
  # @abstract
  class Renderer

    # TODO: move file locations to constants

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
      clean_up()
      template = read_template(self.class.template)
      style = read_style(self.class.style)
      articles = save_images(articles)
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

    # Clean up old berliner.html and berliner_files folder
    # @return [void]
    def clean_up
      FileUtils.rm_rf(File.join(CONFIG_DIR, "berliner_files"))
      FileUtils.rm_rf(File.join(CONFIG_DIR, "berliner.html"))
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

    # Save all images in articles to disk and replace URL with relative filename
    # @param [Array<Article>] articles an array of {Article} objects
    # @return [Array<Article>] an array of {Article} objects with image attribute altered
    def save_images(articles)
      Parallel.map(articles, :in_threads=>10) do |article|
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

    # Download an image given its url, save to disk, and return relative file location
    # @param [String, nil] url the image url
    # @return [String, nil] relative file location of image on disk, or nil if no image
    def save_image(url)
      return nil unless url
      uri = URI.parse(url)
      files_dir = File.join(CONFIG_DIR, "berliner_files")
      unless File.directory?(files_dir)
        FileUtils.mkdir_p(files_dir)
      end
      basename = File.basename(uri.path)
      file = File.join(CONFIG_DIR, "berliner_files", basename)
      begin
        Timeout::timeout(10) {
          File.open(file, 'wb') {|f| f.write(open(uri).read)}
        }
        relative = Pathname.new(file).relative_path_from(Pathname.new(CONFIG_DIR)).to_s
        return relative
      rescue
        return nil
      end
    end

  end
end