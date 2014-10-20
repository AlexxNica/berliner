require "yaml"
require "broadsheet/config"
require "broadsheet/article"

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

    renderer = @profile["renderer"]
    require "broadsheet/renderers/#{renderer}"

    begin
      renderer_klass = renderer.classify.constantize
    rescue
      raise NameError,
            "Error: The #{renderer} renderer was not found. " \
            "Make sure it is defined in renderers/#{renderer}.rb."
    end

    @articles = @sources.map{ |source| source.articles }.flatten
    renderer_klass.render(@articles) # options will get passed in here
  end

end

