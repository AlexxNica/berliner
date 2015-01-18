require "active_support"
require "active_support/core_ext"

module Berliner
  # Manages all Berliner filters
  class FilterManager
    # List all user-defined filters and packaged filters
    # @note Unlike {SourceManager.search}, {search} does not take
    #   a query argument (as there are many less total filters).
    # @return [Array<String>] the slugs of all filters
    def search
      user_filters = Dir[File.join(Dir.home, ".berliner/filters/*")]
      gem_filters = Dir[File.join(LIB_DIR, "berliner/filters/*")]
      filter_slugs = (user_filters + gem_filters).map do |path|
        filename = File.basename(path, ".rb")
        filename.gsub(/_/, "-")
      end
      filter_slugs.uniq.sort
    end

    # Load an instantiated {Source} object(s) given the source slug(s)
    # @param [String, Array<String>] argv the source slug or an array of source slugs
    # @return [Source, Array<Source>] an instance of the specified source or
    #   an array of instances
    def load(argv)
      if argv.is_a?(Array)
        return argv.map{ |s| get_klass(s)}
      end
      get_klass(argv)
    end

    private

    # Return an instantiated {filter} object given the filter slug
    # @param [String] argv the filter slug
    # @raise [LoadError] if the filter can't be loaded
    # @raise [NameError] if the filter's class name can't be found
    # @return [Source] an instance of the specified filter
    def get_klass(argv)
      args = argv.split(" ")
      slug = args.shift
      filename = slug.deslugify + "_filter"
      begin
        require File.join(Dir.home, ".berliner/filters", filename)
      rescue LoadError
        begin
          require File.join("berliner/filters", filename)
        rescue LoadError
        end
      end
      begin
        klass = "Berliner::#{filename.camelize}".constantize
      rescue
        raise NameError,
          "The #{filename.camelize} was not found. " \
          "Make sure it is defined in filters/#{filename}.rb"
      end
      klass.new(args)
    end

  end
end
