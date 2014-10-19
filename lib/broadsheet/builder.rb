require "yaml"
require "broadsheet/config"
require "broadsheet/article"
require "broadsheet/pdf_writer"

class Builder

  class ProfileError < StandardError
  end

  def initialize
    begin
      @profile = YAML.load_file(Broadsheet::PROFILE_PATH)
    rescue
      raise ProfileError, "Profile unreadable"
    end
  end

  def build
    @sources = @profile["sources"].map do |source_slug|
      require "broadsheet/sources/#{source_slug}"
      source_slug.classify.constantize
    end
    @articles = @sources.map{ |source| source.articles }.flatten
    #PDFWriter.write(@articles, "output.pdf")
  end

end

class String

  def classify
    self.split("-").collect(&:capitalize).join
  end

  def constantize
    Object.const_get(self)
  end

end