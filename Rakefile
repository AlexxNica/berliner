require "bundler/gem_tasks"
require "rspec/core/rake_task"
require "yard"

RSpec::Core::RakeTask.new(:spec) do |r|
  r.verbose = false
  r.rspec_opts = "-c"
end

YARD::Rake::YardocTask.new(:yard) do |task|
  task.files   = ["lib/**/*.rb", "-", "*.md", "LICENSE"]
  task.options = [
    "--output-dir", "doc",
    "--markup", "markdown",
    "--exclude", "lib/berliner/sources/*",
    "--exclude", "lib/berliner/renderers/*"
  ]
end

task :undoc do
  exec "yard stats --list-undoc"
end

task :default => :spec
task :test => :spec
task :doc => :yard

task :console do
  exec "irb -r berliner -I ./lib"
end