require "uri"

module Berliner
  # Recognizes canonical sources, based on a manually updated registry
  # @note The benefit of making the source registry a static, manually
  #   updated dictionary mapping is that we don't have to load all sources
  #   (an incredibly expensive operation) in order to recognize the correct
  #   source from an article URL.
  class SourceRegistry

    # A manually updated dictionary mapping
    # URL fragments to canonical source slugs
    REGISTRY = {
      "disegnodaily.com" => "disegno-daily",
      "fivethirtyeight.com" => "five-thirty-eight",
      "nytimes.com" => "new-york-times",
      "newyorker.com" => "new-yorker",
      "theparisreview.org" => "paris-review",
      "slate.com" => "slate",
      "vox.com" => "vox"
    }

    # For a given URL, "recognize" the correct source slug
    # @param [String] permalink the article URL
    # @return [String, nil] the slug representing the article's
    #   canonical source or nil if not recognized
    def self.get_slug_from_url(permalink)
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