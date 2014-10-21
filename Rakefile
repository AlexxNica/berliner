require "bundler/gem_tasks"
require "rspec/core/rake_task"

RSpec::Core::RakeTask.new(:spec) do |r|
  r.verbose = false
  r.rspec_opts = "-c"
end

task :default => :spec
task :test => :spec

task :console do
  exec "irb -r broadsheet -I ./lib"
end