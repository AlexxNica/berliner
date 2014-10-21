module Broadsheet
  def self.read
    profile = Profile.new
    sources = SourceManager.load(profile.sources)
    renderer = RendererManager.load(profile.renderer)

    articles = sources.map{ |source| source.articles }.flatten
    renderer.render(articles)
  end
end

require "broadsheet/extend/module"
require "broadsheet/config"
require "broadsheet/version"
require "broadsheet/source_manager"
require "broadsheet/renderer_manager"
require "broadsheet/profile"