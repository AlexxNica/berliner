require "active_support"
require "active_support/core_ext"
require "berliner/extend/string"

module Berliner
  class SourceManager

    def self.search(foo=nil)
      user_sources = Dir["#{Dir.home}/.berliner/sources/*"]
      gem_sources = Dir["#{LIB_PATH}/berliner/sources/*"]
      source_slugs = (user_sources + gem_sources).map do |path|
        filename = File.basename(path, ".rb")
        filename.gsub(/_/, "-")
      end
      results = source_slugs.uniq.sort
      results = results.grep(foo.query_regex) if foo
      results
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
        require "#{Dir.home}/.berliner/sources/#{filename}"
      rescue LoadError
        require "berliner/sources/#{filename}"
      rescue
      end
      begin
        klass = filename.classify.constantize
      rescue
        raise NameError,
          "The #{filename.classify} source was not found. " \
          "Make sure it is defined in sources/#{filename}.rb"
      end
      klass.new
    end

  end
end