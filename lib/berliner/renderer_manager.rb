require "active_support"
require "active_support/core_ext"

module Berliner
  # Manages all Berliner renderers
  class RendererManager
    # List all user-defined renderers and packaged renderers
    # @note Unlike {SourceManager.search}, {search} does not take
    #   a query argument (as there are many less total renderers).
    # @return [Array<String>] the slugs of all renderers
    def self.search
      user_renderers = Dir["#{Dir.home}/.berliner/renderers/*"]
      gem_renderers = Dir["#{LIB_PATH}/berliner/renderers/*"]
      renderer_slugs = (user_renderers + gem_renderers).map do |path|
        filename = File.basename(path, ".rb")
        filename.dasherize
      end
      renderer_slugs.uniq.sort
    end

    # Load an instantiated {Renderer} object(s) given the renderer slug(s)
    # @param [String, Array<String>] slug the renderer slug or an array of renderer slugs
    # @return [Source, Array<Source>] an instance of the specified renderer or
    #   an array of instances
    def self.load(slug)
      get_klass(slug)
    end

    private
    
    # Return an instantiated {Renderer} object given the renderer slug
    # @param [String] slug the renderer slug
    # @raise [LoadError] if the renderer can't be loaded
    # @raise [NameError] if the renderer's class name can't be found 
    # @return [Source] an instance of the specified renderer
    def self.get_klass(slug)
      filename = slug.gsub(/-/, "_") + "_renderer"
      begin
        require "#{Dir.home}/.berliner/renderers/#{filename}"
      rescue LoadError
        require "berliner/renderers/#{filename}"
      rescue
      end
      begin
        klass = "Berliner::#{filename.classify}".constantize
      rescue
        raise NameError,
          "The #{filename.classify} was not found. " \
          "Make sure it is defined in renderers/#{filename}.rb"
      end
      klass.new
    end

  end
end