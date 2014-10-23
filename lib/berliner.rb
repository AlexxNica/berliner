# Daily digest of online news in a beautiful format
module Berliner
  # Generate and render a Berliner based on the profile
  # @return [void]
  def self.read
    profile = Profile.new
    sources = SourceManager.load(profile.sources)
    renderer = RendererManager.load(profile.renderer)
    articles = sources.map{ |source| source.articles }.flatten
    renderer.render(articles)
  end

  # Search all sources for query term or list all sources if no query is given
  # @param [String, Regexp, nil] foo the search term
  # @return [Array<String>] the sources with foo in their slugs or all sources
  def self.search(foo)
    SourceManager.search(foo)
  end

  # Add a source or sources to the profile if valid
  # @param [String, Array<String>] source the source(s) to add
  # @return [void]
  def self.add(source)
    Profile.new.add(source)
  end

  # Remove a source or sources from the profile
  # @param [String, Array<String>] source the source(s) to remove
  # @return [void]
  def self.remove(source)
    Profile.new.remove(source)
  end

  # List the sources currently added to the profile
  # @return [Array<String>] sources in the profile
  def self.list
    Profile.new.sources
  end

end

require "berliner/extend/module"
require "berliner/config"
require "berliner/version"
require "berliner/source_manager"
require "berliner/renderer_manager"
require "berliner/profile"