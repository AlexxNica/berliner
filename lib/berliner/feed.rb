require "parallel"
require "berliner/source_manager"

module Berliner
  # An Enumerable object which represents the collective feed of all given sources
  class Feed
    attr_accessor :sources, :entries, :articles

    # Create a new {Feed} object. Note that the sources in a {Feed] cannot be
    # changed after it is initialized. This is to allow reliable caching of the
    # entries list, which is only fetched once and then afterwards served from
    # cache. (A more sophisticated caching mechanism could allow the
    # modification of sources post-init, but seems not particularly useful ATM.)
    def initialize(sources)
      @sources = sources
      @entries = get_entries()
    end

    # @return [Array<Article>] a list of articles in this feed
    def articles
      Parallel.map(entries, :in_threads=>10) do |entry|
        entry.article
      end.compact
    end

    private

    # Fetches RSS feed entries for each source
    # @return [Array<Feed::FeedEntry>] a list of entries in this feed
    def get_entries
      Parallel.map(sources, :in_threads=>10) do |source|
        source.fetch
      end.flatten.compact
    end

    # A single entry to a Berliner {Feed}.
    # A {FeedEntry} stores metadata about an article such as the URL.
    # It also scrapes article content when requested, and caches that content.
    class FeedEntry
      attr_accessor :url, :via, :source

      # Create a new {Feed::FeedEntry} object
      def initialize(url, via, source = nil)
        @url = url
        @via = via
        @source = source || SourceManager.load_from_url(url)
      end

      # @return [Article] the article associated with this entry
      def article
        @article ||= get_article
      end

      private

      def get_article
        begin
          source.parse(self)
        rescue
          nil
        end
      end
    end

  end
end
