require "uri"

module Berliner
  class SourceRegistry

    REGISTRY = {
      "disegnodaily.com" => "disegno-daily",
      "fivethirtyeight.com" => "five-thirty-eight",
      "nytimes.com" => "new-york-times",
      "newyorker.com" => "new-yorker",
      "theparisreview.org" => "paris-review",
      "slate.com" => "slate",
      "vox.com" => "vox"
    }

    def self.get_classname(permalink)
      begin
        uri = URI(permalink)
        REGISTRY.each do |domain, slug|
          return slug if uri.host.include? domain
        end
      rescue
        return nil
      end
      return nil
    end
  end
end