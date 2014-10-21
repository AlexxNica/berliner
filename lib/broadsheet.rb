module Broadsheet
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

require "broadsheet/extend/module"
require "broadsheet/config"
require "broadsheet/version"
require "broadsheet/source_manager"
require "broadsheet/renderer_manager"
require "broadsheet/profile"