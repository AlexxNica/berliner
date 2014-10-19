# -*- encoding: utf-8 -*-
lib = File.expand_path("../lib", __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require "broadsheet/config"

Gem::Specification.new do |gem|
  gem.name          = "broadsheet"
  gem.version       = Broadsheet::VERSION
  gem.authors       = ["Seth Thompson", "Geoffrey Litt"]
  gem.email         = ["s3th.thompson@gmail.com", "gklitt@gmail.com"]
  gem.description   = "Daily digest of online news in a beautiful format."
  gem.summary       = "Broadsheet daily digest"
  gem.homepage      = "https://github.com/s3ththompson/broadsheet"
  gem.licenses      = ["MIT"]

  gem.files         = `git ls-files`.split($/)
  gem.executables   = gem.files.grep(%r{^bin/}).map{ |f| File.basename(f) }
  gem.test_files    = gem.files.grep(%r{^(test|spec|features)/})
  gem.require_paths = ["lib"]

  gem.add_runtime_dependency "feedjira", ["= 1.1.0"]
  gem.add_runtime_dependency "prawn", ["= 1.3.0"]
  gem.add_runtime_dependency "ruby-readability", ["= 0.7.0"]
end