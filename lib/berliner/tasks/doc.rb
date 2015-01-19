require "yard"

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

task :doc => :yard