module Berliner
  # Description of the gem
  DESCRIPTION = "Daily digest of online news in a beautiful format"

  # Path to the config directory
  CONFIG_DIR = ENV.fetch("BERLINER_CONFIG", File.join(Dir.home, ".berliner"))

  # Path to the profile in the config directory
  PROFILE_PATH = File.join(CONFIG_DIR, "profile.yml")
  # Path to gem lib/ folder
  LIB_DIR = File.expand_path("..", __dir__)
end
