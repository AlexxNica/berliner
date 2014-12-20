require 'berliner/source_manager'
require 'parallel'

module Berliner
  # An Enumerable object which represents the collective feed of all given sources
  class Feed
    include Enumerable
    attr_accessor :sources

    # Create a new {Feed} object
    def initialize(sources)
      @sources = sources
    end

    # Enumerator which constitutes an infinite stream of articles
    # Sources are downloaded and in parallel
    # A {Feed} object should be lazily evaluated for best performance
    # @yield [Article] Gives the next {Article} in the feed to the block
    def each
      entries = Parallel.map(sources, :in_threads=>10){|source| source.fetch }
      # interleave source entries
      first, *rest = *entries
      entries = first.zip(*rest).flatten.compact
      Parallel.each(entries, :in_threads=>10){ |entry| yield parse(entry) }
    end

    private

    # Parse a {Feed::FeedEntry} into an {Article}
    # @param [Feed::FeedEntry] entry the feed entry
    # @return [Article] the full, parsed article
    def parse(entry)
      source = SourceManager.load_from_url(entry.url)
      source.parse(entry)
    end

    # A single entry to a Berliner {Feed}
    class FeedEntry
      attr_accessor :url, :via

      # Create a new {Feed::FeedEntry} object
      def initialize(url, via)
        @url = url
        @via = via
      end
    end

  end
end