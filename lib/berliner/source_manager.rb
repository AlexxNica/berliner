require "active_support"
require "active_support/core_ext"
require "berliner/extend/string"
require "berliner/source"
require "berliner/source_registry"
require "berliner/loader"

module Berliner
  # Manages all Berliner sources
  class SourceManager
    # Create a new {SourceManager} object
    # @param [Hash, nil] credentials a dictionary with
    #   slugs as keys to actual creds (sub-)dictionaries, or nil
    # @note "credentials" is used as the plural to refer to more than one
    #   source's authentication tokens.  "creds" is used as the plural to
    #   refer to the many tokens that a single source might require for
    #   authorization.
    def initialize(credentials = {})
      @credentials = credentials
    end

    # Search user-defined sources and packaged sources for a query term foo
    # or list all sources if foo is nil.
    # @param [String, Regex, nil] foo the query term
    # @return [Array<String>] the slugs of all sources with foo in their slugs
    #   or all sources
    def search(foo = nil)
      user_sources, gem_sources = Loader.list_files("berliner/sources/")
      source_slugs = (user_sources + gem_sources).map do |path|
        filename = File.basename(path, ".rb")
        filename.slugify
      end
      results = source_slugs.uniq.sort
      results = results.grep(foo.query_regex) if foo
      results
    end

    # Load an instantiated {Source} object(s) given the source slug(s)
    # @param [String, Array<String>] slug the source slug or an array of source
    #   slugs
    # @return [Source, Array<Source>] an instance of the specified source or
    #   an array of instances
    def load(slug)
      if slug.is_a?(Array)
        return slug.map do |s|
          get_klass(s)
        end
      end
      get_klass(slug)
    end

    # Load an instantiated {Source} object(s) given an article permalink(s)
    # @param [String, Array<String>] permalink the article permalink or an array
    #   of permalinks
    # @return [Source, Array<Source>] an instance of the specified source or
    #   an array of instances
    def load_from_url(permalink)
      if permalink.is_a?(Array)
        return permalink.map { |s| get_klass_from_url(s) }
      end
      get_klass_from_url(permalink)
    end

    private

    # Use a class instance variable to implement a cache of loaded classes
    # across all instances of {SourceManager}
    # Useful in order to retrieve previously credentialed {Source} instance
    # without credentials
    # @return [Hash] a cache of classes
    def self.klasses
      @klasses ||= {}
    end

    # Return an instantiated {Source} object given the source slug
    # @param [String] slug the source slug
    # @raise [LoadError] if the source can't be loaded
    # @raise [NameError] if the source's class name can't be found
    # @return [Source] an instance of the specified source
    def get_klass(slug)
      # Check cache for slug
      if self.class.klasses.key?(slug)
        # Unless credentials are provided and cached {Source} instance
        # was uncredentialed, return the cached instance
        unless get_creds(slug) && !self.class.klasses[slug].authenticated
          return self.class.klasses[slug]
        end
      end
      filename = slug.deslugify
      begin
        klass = Loader.read_klass(File.join("berliner", "sources", filename))
      rescue
        raise NameError,
              "The #{filename.camelize} source was not found. " \
              "Make sure it is defined in sources/#{filename}.rb"
      end
      creds = get_creds(slug)
      k = klass.new(creds)
      self.class.klasses[slug] = k
      k
    end

    # Return an instantiated {Source} object given an article permalink
    # @param [String] permalink the article permalink
    # @return [Source, DefaultSource] an instance of the recognized source
    #   or the default source
    def get_klass_from_url(permalink)
      slug = SourceRegistry.get_slug_from_url(permalink)
      if slug
        return get_klass(slug)
      else
        return DefaultSource.new
      end
    end

    # Return the source's creds (from the credentials dictionary)
    # @param [String] slug the source slug
    # @return [Hash, nil] the creds hash or nil
    # @note See {#initialize} for the difference between "credentials"
    #   and "creds"
    def get_creds(slug)
      @credentials[slug] || nil
    end
  end
end
