lib = File.expand_path("../lib", __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)

require "berliner/tasks/doc"
require "berliner/tasks/gem"
require "berliner/tasks/sources"
require "berliner/tasks/test"

task :console do
  exec "irb -r berliner -I ./lib"
end
