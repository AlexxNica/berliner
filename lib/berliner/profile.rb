require "fileutils"
require "yaml"
require "active_support"
require "active_support/core_ext"

module Berliner
  # A Berliner profile stores per-user preferences and configuration
  class Profile
    attr_accessor :profile

    PROFILE_PATH = File.join(CONFIG_DIR, "profile.yml")

    # Create a new {Profile} object
    def initialize
      begin
        FileUtils.mkdir_p(File.dirname(PROFILE_PATH))
        user_profile = YAML.load_file(PROFILE_PATH)
        @profile = default_profile.merge(user_profile)
      rescue
        @profile = default_profile
      end
    end

    # Add a source to the profile if the source is valid
    # @param [String, Array<String>] source a source slug or
    #   array of source slugs
    # @raise [NameError] if the source is not found
    # @return [void]
    def add(source)
      if source.is_a?(Array)
        source.each{ |s| add(s)}
        return true
      end
      if SourceManager.search.include?(source)
        profile["sources"] |= [source]
      else
        raise NameError, "Source #{source} not found"
      end
      write
    end

    # Remove a source from the profile
    # @param [String, Array<String>] source a source slug or
    #   array of source slugs
    # @return [void]
    def remove(source)
      if source.is_a?(Array)
        source.each{ |s| remove(s)}
        return true
      end
      profile["sources"] -= [source]
      write
    end

    # List the sources saved in the profile
    # @return [Array<String>] an array of source slugs
    def sources
      profile["sources"]
    end

    # List the renderer saved in the profile
    # @return [<String>] a renderer slug
    def renderer
      profile["renderer"] || "default"
    end

    # List the filters saved in the profile
    # @return [Array<String>] an array of filter slugs
    def filters
      profile["filters"]
    end

    def credentials
      profile["credentials"] || {}
    end

    private

    # Default profile object
    # @return [Hash] the default profile object
    def default_profile
      {
        "sources" => [],
        "filters" => [],
        "credentials" => {},
        "renderer" => "default"
      }
    end

    # Write the profile to disk
    # @return [void]
    def write
      File.open(PROFILE_PATH,"w") do |f|
         f.write profile.to_hash.to_yaml
      end
    end

  end
end
