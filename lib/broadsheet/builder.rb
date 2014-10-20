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
    @sources = @profile["sources"].map do |source|
      filename = source.gsub(/-/, '_')
      require "broadsheet/sources/#{filename}"
      begin
        source_klass = filename.classify.constantize
      rescue
        raise NameError,
              "Error: The #{filename.classify} source was not found. " \
              "Make sure it is defined in sources/#{filename}.rb."
      end
      source_klass.new
    end
  end

  def load_renderer
    renderer = @profile["renderer"]
    filename = renderer.gsub(/-/, '_')
    require "broadsheet/renderers/#{renderer}"
    begin
      renderer_klass = filename.classify.constantize
    rescue
      raise NameError,
            "Error: The #{filename.classify} renderer was not found. " \
            "Make sure it is defined in renderers/#{filename}.rb."
    end

    renderer_klass.new  # when we have an options hash, pass it in here
  end

end

