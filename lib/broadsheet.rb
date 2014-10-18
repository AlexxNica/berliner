module Broadsheet
  def self.hi
    puts Builder.new.build
  end
end

require 'broadsheet/extend/module'
require 'broadsheet/version'
require 'broadsheet/builder'