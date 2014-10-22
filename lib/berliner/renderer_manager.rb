require "active_support"
require "active_support/core_ext"

module Berliner
  class RendererManager

    def self.search
      user_renderers = Dir["#{Dir.home}/.berliner/renderers/*"]
      gem_renderers = Dir["#{LIB_PATH}/berliner/renderers/*"]
      renderer_slugs = (user_renderers + gem_renderers).map do |path|
        filename = File.basename(path, ".rb")
        filename.gsub(/_/, "-")
      end
      renderer_slugs.uniq.sort
    end

    def self.load(slug)
      get_klass(slug)
    end

    private

    def self.get_klass(slug)
      filename = slug.gsub(/-/, "_") + "_renderer"
      begin
        require "#{Dir.home}/.berliner/renderers/#{filename}"
      rescue LoadError
        require "berliner/renderers/#{filename}"
      rescue
      end
      begin
        klass = filename.classify.constantize
      rescue
        raise NameError,
          "The #{filename.classify} was not found. " \
          "Make sure it is defined in renderers/#{filename}.rb"
      end
      klass.new
    end

  end
end