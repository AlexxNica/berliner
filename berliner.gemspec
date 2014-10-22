# -*- encoding: utf-8 -*-
lib = File.expand_path("../lib", __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require "berliner/version"

Gem::Specification.new do |gem|
  gem.name          = "berliner"
  gem.version       = Berliner::VERSION
  gem.authors       = ["Seth Thompson", "Geoffrey Litt"]
  gem.email         = ["s3th.thompson@gmail.com", "gklitt@gmail.com"]
  gem.description   = "Berliner is a Ruby gem and CLI that compiles a daily digest of online news in a beautiful format."
  gem.summary       = "Daily news digest."
  gem.homepage      = "https://github.com/s3ththompson/berliner"
  gem.licenses      = ["MIT"]

  gem.files         = `git ls-files`.split($/)
  gem.executables   = ["berliner"]
  gem.test_files    = gem.files.grep(%r{^(test|spec|features)/})
  gem.require_paths = ["lib"]

  gem.required_ruby_version = ">= 2.0.0"

  gem.add_development_dependency "rake", ["= 10.1.0"]
  gem.add_development_dependency "rspec", ["= 3.1.0"]
  gem.add_development_dependency "coveralls", ["= 0.7.1"]

  gem.add_runtime_dependency "feedjira", ["= 1.1.0"]
  gem.add_runtime_dependency "prawn", ["= 1.3.0"]
  gem.add_runtime_dependency "ruby-readability", ["= 0.7.0"]
  gem.add_runtime_dependency "activesupport", ["= 4.1.6"]
  gem.add_runtime_dependency "nokogiri", ["= 1.6.3.1"]
  gem.add_runtime_dependency "commander", ["= 4.2.1"]
end
