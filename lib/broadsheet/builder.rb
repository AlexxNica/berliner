require "yaml"
require "broadsheet/config"
require "broadsheet/source_manager"
require "broadsheet/renderer_manager"

class Builder

  class ProfileError < StandardError
  end

  def initialize
    begin
      @profile = YAML.load_file(Broadsheet::PROFILE_PATH)
    rescue
      raise ProfileError, "Profile unreadable"
    end

    raise "Error: profile.yaml must contain a sources key." unless @profile["sources"]
    raise "Error: profile.yaml must contain a renderer key." unless @profile["renderer"]
  end

  def build
    sources = SourceManager.load(@profile["sources"])
    renderer = RendererManager.load(@profile["renderer"])

    articles = sources.map{ |source| source.articles }.flatten
    renderer.render(articles)
  end

end

