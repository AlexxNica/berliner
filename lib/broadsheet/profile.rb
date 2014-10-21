require "yaml"
require 'active_support/core_ext/hash'

module Broadsheet
  class Profile

    def initialize
      begin
        @profile = YAML.load_file(Broadsheet::PROFILE_PATH).with_indifferent_access
      rescue
        @profile = {
          sources: "test",
          renderer: "console"
        }
      end
    end

    def sources
      @profile[:sources]
    end

    def renderer
      @profile[:renderer]
    end

  end
end