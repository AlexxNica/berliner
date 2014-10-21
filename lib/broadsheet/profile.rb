require "yaml"
require 'active_support/core_ext/hash'

module Broadsheet
  class Profile

    def initialize
      begin
        @profile = YAML.load_file(Broadsheet::PROFILE_PATH).with_indifferent_access
      rescue
        @profile = {
          sources: [],
          renderer: "console"
        }
      end
    end

    def add(source)
      if source.is_a?(Array)
        source.each{ |s| add(s)}
        return true
      end
      @profile[:sources] |= [source]
      write
    end

    def remove(source)
      if source.is_a?(Array)
        source.each{ |s| remove(s)}
        return true
      end
      @profile[:sources] -= [source]
      write
    end

    def sources
      @profile[:sources]
    end

    def renderer
      @profile[:renderer]
    end

    private

    def write
      File.open(Broadsheet::PROFILE_PATH,'w') do |f| 
         f.write @profile.to_yaml
      end
    end

  end
end