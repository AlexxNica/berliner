require "active_support"
require "active_support/core_ext"
require "berliner/extend/string"
require "berliner/source"

module Berliner
  # Manages all Berliner sources
  class SourceManager
    # Search user-defined sources and packaged sources for a query term foo
    # or list all sources if foo is nil.
    # @param [String, Regex, nil] foo the query term
    # @return [Array<String>] the slugs of all sources with foo in their slugs or all sources
    def self.search(foo=nil)
      user_sources = Dir[File.join(Dir.home, ".berliner/sources/*")]
      gem_sources = Dir[File.join(LIB_DIR, "berliner/sources/*")]
      source_slugs = (user_sources + gem_sources).map do |path|
        filename = File.basename(path, ".rb")
        filename.dasherize
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

    # Load an instantiated {Source} object(s) given an article permalink(s)
    # @param [String, Array<String>] permalink the article permalink or an array of permalinks
    # @return [Source, Array<Source>] an instance of the specified source or
    #   an array of instances
    def self.load_from_url(permalink)
      if permalink.is_a?(Array)
        return permalink.map{ |s| get_klass_from_url(s)}
      end
      get_klass_from_url(permalink)
    end

    private

    # Return an instantiated {Source} object given the source slug
    # @param [String] slug the source slug
    # @raise [LoadError] if the source can't be loaded
    # @raise [NameError] if the source's class name can't be found
    # @return [Source] an instance of the specified source
    def self.get_klass(slug)
      if @all_klasses
        found = @all_klasses.find{|source| slug == source.class.title}
        return found if found
      end

      filename = slug.gsub(/-/, "_")
      begin
        require File.join(Dir.home, ".berliner/sources", filename)
      rescue LoadError
        begin
          require File.join("berliner/sources", filename)
        rescue LoadError
        end
      end
      begin
        klass = "Berliner::#{filename.classify}".constantize
      rescue
        raise NameError,
          "The #{filename.classify} source was not found. " \
          "Make sure it is defined in sources/#{filename}.rb"
      end
      klass.new
    end

    # Return an instantiated {Source} object given an article permalink
    # @param [String] permalink the article permalink
    # @return [Source, DefaultSource] an instance of the recognized source or the default source
    def self.get_klass_from_url(permalink)
      found = self.get_all_klasses().find{|source| source.recognize?(permalink)}
      return found || DefaultSource.new
    end

    # Return all instantiated {Source} objects (memoized)
    # @return [Array<Source>] instances of all sources
    def self.get_all_klasses()
      @all_klasses ||= self.expensive_all_klasses()
    end

    # Return all instantiated {Source} objects
    # @return [Array<Source>] instances of all sources
    def self.expensive_all_klasses()
      self.search().map{ |s| get_klass(s) }
    end

  end
end
