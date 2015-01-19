require "parallel"
require "httpclient"
require "berliner/source_manager"

module Berliner
  # An Enumerable object which represents the collective feed of all given sources
  class Feed
    attr_accessor :s_manager, :sources, :entries, :articles

    # Create a new {Feed} object. Note that the sources in a {Feed] cannot be
    # changed after it is initialized. This is to allow reliable caching of the
    # entries list, which is only fetched once and then afterwards served from
    # cache. (A more sophisticated caching mechanism could allow the
    # modification of sources post-init, but seems not particularly useful ATM.)
    def initialize(slugs, credentials = {})
      @s_manager = SourceManager.new(credentials)
      @sources = s_manager.load(slugs)
      @entries = get_entries
    end

    # @return [Array<Article>] a list of articles in this feed
    def articles
      Parallel.map(entries, in_threads: 10) do |entry|
        entry.article
      end.compact
    end

    private

    # Fetches RSS feed entries for each source
    # @return [Array<Feed::FeedEntry>] a list of entries in this feed
    def get_entries
      Parallel.map(sources, in_threads: 10) do |source|
        source.fetch.map do |url|
          s = s_manager.load_from_url(url)
          FeedEntry.new(url, source.class.title, s)
        end
      end.flatten.compact
    end

    # A single entry to a Berliner {Feed}.
    # A {FeedEntry} stores metadata about an article such as the URL.
    # It also scrapes article content when requested, and caches that content.
    class FeedEntry
      attr_accessor :url, :via, :source

      # Create a new {Feed::FeedEntry} object
      # The {Source} instance for the feed is designed to be passed in from
      # the parent {Feed} instance so that the sources are properly
      # credentialed.  However, an (uncredentialed) {Source} instance will be
      # created for the given URL, if not passed in.
      def initialize(rss_url, via, source = nil)
        @url = follow_url(rss_url)
        @via = via
        @source = source || SourceManager.new.load_from_url(url)
      end

      # @return [Article] the article associated with this entry
      def article
        @article ||= get_article
      end

      private

      # Follows redirects on the article URL in the feed entry
      # to arrive at the canonical article URL
      # @example Follow NYT RSS URL to canonical article permalink
      #   follow_url("http://rss.nytimes.com/c/34625/f/640350/s/427271db/sc/8/l/0L0Snytimes0N0C20A150C0A10C190Cnyregion0Cicy0Eroads0Ecause0Eaccidents0Eand0Eclose0Ebridges0Bhtml0Dpartner0Frss0Gemc0Frss/story01.htm") #=>
      #     "http://www.nytimes.com/2015/01/19/nyregion/icy-roads-cause-accidents-and-close-bridges.html?partner=rss&emc=rss"
      # @param [String] rss_url an article URL (as found in a feed entry)
      # @return [String] a canonical article URL
      def follow_url(rss_url)
        httpc = HTTPClient.new
        resp = httpc.get(rss_url)
        # If no response header, just return the original URL
        location = resp.header["Location"] || rss_url
        # Get the first URL in the Location header
        location = location.shift if location.is_a?(Array)
        # Sanity check that a URL was given
        (location && !location.empty?) ? location : rss_url
      rescue
        nil
      end

      # Parse the fetched article
      # @return [Article] an article
      def get_article
        source.parse(self)
      rescue
        nil
      end
    end
  end
end
