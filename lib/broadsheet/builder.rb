require "yaml"
require "broadsheet/config"
require "broadsheet/article"
require 'active_support/core_ext'

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
    load_sources
    renderer = load_renderer

    @articles = @sources.map{ |source| source.articles }.flatten
    renderer.render(@articles)
  end

  private

  def load_sources
    @sources = @profile["sources"].map do |source_slug|
      require "broadsheet/sources/#{source_slug}"
      begin
        source_slug.classify.constantize
      rescue
        raise NameError,
              "Error: The #{source_slug.classify} source was not found. " \
              "Make sure it is defined in sources/#{source_slug}.rb."
      end
    end
  end

  def load_renderer
    renderer = @profile["renderer"]
    require "broadsheet/renderers/#{renderer}"

    begin
      renderer_klass = renderer.classify.constantize
    rescue
      raise NameError,
            "Error: The #{renderer} renderer was not found. " \
            "Make sure it is defined in renderers/#{renderer}.rb."
    end

    renderer_klass.new  # when we have an options hash, pass it in here
  end

end

