require "broadsheet/builder"

module Broadsheet
  def self.read
    Builder.new.build
  end
end

require "broadsheet/extend/module"
require "broadsheet/config"
