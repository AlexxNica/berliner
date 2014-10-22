module Berliner
  def self.read
    profile = Profile.new
    sources = SourceManager.load(profile.sources)
    renderer = RendererManager.load(profile.renderer)
    articles = sources.map{ |source| source.articles }.flatten
    renderer.render(articles)
  end

  def self.search(foo)
    SourceManager.search(foo)
  end

  def self.add(source)
    Profile.new.add(source)
  end

  def self.remove(source)
    Profile.new.remove(source)
  end

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