require "berliner/config"
require "berliner/version"
require "berliner/source_manager"
require "berliner/renderer_manager"
require "berliner/filter_manager"
require "berliner/profile"
require "berliner/article"
require "berliner/echo"

# Daily digest of online news in a beautiful format
module Berliner
  # Implements top-level CLI commands
  class CLI
    attr_accessor :profile

    # Create a new {CLI} object
    def initialize(verbose = false)
      Echo.verbosity = verbose ? Echo::VERBOSE : Echo::NORMAL
    end

    # Generate and render a Berliner based on the profile
    # @return [void]
    def read
      @profile = Profile.new
      Echo.info("Fetching articles...")
      feed = Feed.new(profile.sources, profile.credentials)
      renderer = RendererManager.new.load(profile.renderer)
      filters = FilterManager.new.load(profile.filters)
      Echo.info("Filtering articles...")
      feed = filters.inject(feed) do |f, filter|
        filter.filter(f)
      end
      Echo.info("Rendering...")
      output = renderer.render(feed.articles)
      Echo.success("Opening Berliner")
      begin
        system %(open "#{output}")
      rescue
        pass
      end
    end

    # Search all sources for query term or list all sources if no query is given
    # @param [String, Regexp, nil] foo the search term
    # @return [Array<String>] the sources with foo in their slugs or all sources
    def search(foo)
      SourceManager.new.search(foo).each do |source|
        Echo.info(source)
      end
    end

    # Add a source or sources to the profile if valid
    # @param [String, Array<String>] source the source(s) to add
    # @return [void]
    def add(source)
      @profile = Profile.new
      profile.add(source)
    end

    # Remove a source or sources from the profile
    # @param [String, Array<String>] source the source(s) to remove
    # @return [void]
    def remove(source)
      @profile = Profile.new
      profile.remove(source)
    end

    # List the sources currently added to the profile
    # @return [Array<String>] sources in the profile
    def list
      @profile = Profile.new
      profile.sources.each do |source|
        Echo.info(source)
      end
    end
  end
end
