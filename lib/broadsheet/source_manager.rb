require "active_support/core_ext"

class SourceManager

  def self.list
  end

  def self.load(slug)
    if slug.is_a?(Array)
      return slug.map{ |s| get_klass(s)}
    end
    get_klass(slug)
  end

  private

  def self.get_klass(slug)
    filename = slug.gsub(/-/, "_")
    begin
      require "#{Dir.home}/.broadsheet/sources/#{filename}"
    rescue LoadError
      require "broadsheet/sources/#{filename}"
    rescue
    end
    begin
      klass = filename.classify.constantize
    rescue
      raise NameError,
        "Error: The #{filename.classify} source was not found. " \
        "Make sure it is defined in sources/#{filename}.rb."
    end
    klass.new
  end

end