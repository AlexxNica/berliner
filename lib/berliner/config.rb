module Berliner
  # Description of the gem
  DESCRIPTION = "Daily digest of online news in a beautiful format"
  # Path to profile in user's home directory
  PROFILE_PATH = File.join(Dir.home, ".berliner/profile.yml")
  # Path to gem lib/ folder
  LIB_PATH = File.expand_path('..', __dir__)
end
