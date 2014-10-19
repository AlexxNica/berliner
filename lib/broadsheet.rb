module Broadsheet
  def self.hi
    puts Builder.new.build
  end
end

require 'broadsheet/extend/module'
require 'broadsheet/version'

require 'broadsheet/article'
require 'broadsheet/builder'
require 'broadsheet/config'
require 'broadsheet/pdf_writer'
