require "active_support"
require "active_support/core_ext"

module Berliner
  # Lazily load Berliner classes
  class Loader

    # @todo implement
    def self.read_file(path)
    end

    # Read a class from the filesystem and return the evaled namespace
    # @param [String] path a relative pathname, ex: "berliner/sources/new_york_times"
    # @return [Object] the requested class object
    # @raise [NameError] if the class can't be found
    def self.read_klass(path)
      p = self.normalize_path(path)
      self.load_klass(p)
      filename = File.basename(p, ".rb")
      begin
        klass = self.constantize(filename)
      rescue
        raise NameError
      end
      klass
    end

    # @todo implement
    def self.list_files(path)
    end

    private

    # Constantize the fully-qualified classname from filename
    # @param [String] filename the filename, ex: "new_york_times"
    # @return [Object] the constantized class object
    def self.constantize(filename)
      "Berliner::#{filename.camelize}".constantize
    end

    # Normalize the path by removing the '.rb' extension if it exists
    # @param [String] path the path
    # @return [String] the path sans '.rb' extension
    def self.normalize_path(path)
      path.chomp(".rb")
    end

    # Require the klass file at the specified path
    # @param [string] p a normalized relative path, ex: "berliner/sources/new_york_times"
    # @return [void]
    def self.load_klass(p)
      parts = p.split(File::SEPARATOR)
      parts.shift
      user_path = File.join(CONFIG_DIR, parts)
      gem_path = File.join(LIB_DIR, "berliner", parts)
      begin
        require user_path
      rescue LoadError
        begin
          require gem_path
        rescue LoadError
        end
      end
    end
  end
end