module Broadsheet
  def self.read
    Builder.new.build
  end
end

require "broadsheet/extend/module"
require "broadsheet/config"
require "broadsheet/version"
require "broadsheet/source_manager"
require "broadsheet/renderer_manager"
require "broadsheet/builder"