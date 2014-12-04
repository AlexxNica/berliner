require "berliner/source_manager"
require "berliner/renderer_manager"
require "berliner/profile"

# Daily digest of online news in a beautiful format
module Berliner

  # Implements top-level CLI commands
  class CLI
    attr_accessor :profile

    # Create a new {CLI} object
    def initialize(verbose: false)
      @profile = Profile.new
    end

    # Generate and render a Berliner based on the profile
    # @return [void]
    def read
      sources = SourceManager.load(profile.sources)
      renderer = RendererManager.load(profile.renderer)
      articles = sources.map{ |source| source.articles }.flatten
      renderer.render(articles)
    end

    # Search all sources for query term or list all sources if no query is given
    # @param [String, Regexp, nil] foo the search term
    # @return [Array<String>] the sources with foo in their slugs or all sources
    def search(foo)
      SourceManager.search(foo)
    end

    # Add a source or sources to the profile if valid
    # @param [String, Array<String>] source the source(s) to add
    # @return [void]
    def add(source)
      profile.add(source)
    end

    # Remove a source or sources from the profile
    # @param [String, Array<String>] source the source(s) to remove
    # @return [void]
    def remove(source)
      profile.remove(source)
    end

    # List the sources currently added to the profile
    # @return [Array<String>] sources in the profile
    def list
      profile.sources
    end
  end
end

require "berliner/config"
require "berliner/version"
require "berliner/filter"