require "active_support"
require "active_support/core_ext"
require "berliner/extend/string"

module Berliner
  # Manages all Berliner sources
  class SourceManager
    # Search user-defined sources and packaged sources for a query term foo
    # or list all sources if foo is nil.
    # @param [String, Regex, nil] foo the query term
    # @return [Array<String>] the slugs of all sources with foo in their slugs or all sources
    def self.search(foo=nil)
      user_sources = Dir["#{Dir.home}/.berliner/sources/*"]
      gem_sources = Dir["#{LIB_PATH}/berliner/sources/*"]
      source_slugs = (user_sources + gem_sources).map do |path|
        filename = File.basename(path, ".rb")
        filename.gsub(/_/, "-")
      end
      results = source_slugs.uniq.sort
      results = results.grep(foo.query_regex) if foo
      results
    end

    # Load an instantiated {Source} object(s) given the source slug(s)
    # @param [String, Array<String>] slug the source slug or an array of source slugs
    # @return [Source, Array<Source>] an instance of the specified source or
    #   an array of instances
    def self.load(slug)
      if slug.is_a?(Array)
        return slug.map{ |s| get_klass(s)}
      end
      get_klass(slug)
    end

    private

    # Return an instantiated {Source} object given the source slug
    # @param [String] slug the source slug
    # @raise [LoadError] if the source can't be loaded
    # @raise [NameError] if the source's class name can't be found 
    # @return [Source] an instance of the specified source
    def self.get_klass(slug)
      filename = slug.gsub(/-/, "_")
      begin
        require "#{Dir.home}/.berliner/sources/#{filename}"
      rescue LoadError
        require "berliner/sources/#{filename}"
      rescue
      end
      begin
        klass = filename.classify.constantize
      rescue
        raise NameError,
          "The #{filename.classify} source was not found. " \
          "Make sure it is defined in sources/#{filename}.rb"
      end
      klass.new
    end

  end
end