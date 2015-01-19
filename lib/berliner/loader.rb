require "active_support"
require "active_support/core_ext"

module Berliner
  # Lazily load Berliner classes
  class Loader
    # Read the contents of a file (checking in the user config folder first,
    # then the gem lib)
    # @param [String] path a relative pathname, with extension,
    #   ex: "berliner/assets/styles/utilitarian"
    # @return [String] the requested class object
    # @raise [NameError] if the class can't be found
    def self.read_file(path)
      user_path, gem_path = user_gem_paths(path)
      begin
        contents = File.read(user_path)
      rescue
        begin
          contents = File.read(gem_path)
        rescue
          raise NameError
        end
      end
      contents
    end

    # Lazily load a class (checking in the user config folder first,
    # then the gem lib)
    # @param [String] path a relative pathname,
    #   ex: "berliner/sources/new_york_times"
    # @return [Object] the requested class object
    # @raise [NameError] if the class can't be found
    def self.read_klass(path)
      p = normalize_path(path)
      load_klass(p)
      filename = File.basename(p, ".rb")
      begin
        klass = constantize(filename)
      rescue
        raise NameError
      end
      klass
    end

    # List the files in a directory (from the user's config dir and the gem lib)
    # @param [String] path the path to a relative directory,
    #   ex: "berliner/renderers/"
    # @return [Array(Array<String>, Array<String>)] a tuple of arrays of
    #   filepaths the first element in the tuple is the array of filepaths from
    #   the user's config dir, and the second element in the tuple is that of
    #   the gem lib
    def self.list_files(path)
      user_path, gem_path = user_gem_paths(File.join(path, "*"))
      [Dir[user_path], Dir[gem_path]]
    end

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

    # For a given relative path, get the full filepath in the user's
    # config dir, and the gem lib dir
    # @param [String] path the relative path, ex: "berliner/renderers/"
    # return [Array(String, String)] a tuple of the user config filepath and the
    #   gem lib filepath,
    #   ex: ("/home/.berliner/renderers", "/gem/lib/berliner/renderers")
    def self.user_gem_paths(path)
      parts = path.split(File::SEPARATOR)
      parts.shift
      user_path = File.join(CONFIG_DIR, parts)
      gem_path = File.join(LIB_DIR, "berliner", parts)
      [user_path, gem_path]
    end

    # Require the klass file at the specified path
    # @param [string] p a normalized relative path,
    #   ex: "berliner/sources/new_york_times"
    # @return [void]
    def self.load_klass(p)
      user_path, gem_path = user_gem_paths(p)
      begin
        require user_path
      rescue LoadError
        begin
          require gem_path
        rescue LoadError
          pass
        end
      end
    end
  end
end
