require "active_support"
require "active_support/core_ext"

module Berliner
  # Manages all Berliner filters
  class FilterManager
    # List all user-defined filters and packaged filters
    # @note Unlike {SourceManager.search}, {search} does not take
    #   a query argument (as there are many less total filters).
    # @return [Array<String>] the slugs of all filters
    def self.search
      user_filters = Dir["#{Dir.home}/.berliner/filters/*"]
      gem_filters = Dir["#{LIB_PATH}/berliner/filters/*"]
      filter_slugs = (user_filters + gem_filters).map do |path|
        filename = File.basename(path, ".rb")
        filename.gsub(/_/, "-")
      end
      filter_slugs.uniq.sort
    end

    # Load an instantiated {Source} object(s) given the source slug(s)
    # @param [String, Array<String>] slug the source slug or an array of source slugs
    # @return [Source, Array<Source>] an instance of the specified source or
    #   an array of instances
    def self.load(slug)
      if slug.is_a?(Array)
        return slug.map{ |s| get_klass(s)}
      end
      get_klass(slug)
    end

    private

    # Return an instantiated {filter} object given the filter slug
    # @param [String] slug the filter slug
    # @raise [LoadError] if the filter can't be loaded
    # @raise [NameError] if the filter's class name can't be found
    # @return [Source] an instance of the specified filter
    def self.get_klass(slug)
      filename = slug.gsub(/-/, "_") + "_filter"
      begin
        require "#{Dir.home}/.berliner/filters/#{filename}"
      rescue LoadError
        require "berliner/filters/#{filename}"
      rescue
      end
      begin
        klass = "Berliner::#{filename.classify}".constantize
      rescue
        raise NameError,
          "The #{filename.classify} was not found. " \
          "Make sure it is defined in filters/#{filename}.rb"
      end
      klass.new
    end

  end
end
