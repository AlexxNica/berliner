require "feedjira"
require "ruby-readability"
require "open-uri"
require "open_uri_redirections"
require "berliner/article"
require "berliner/extend/string"
require "berliner/extend/module"
require "berliner/feed"
require "active_support"
require "active_support/core_ext"
require "uri"

module Berliner
  # The base object for a Berliner source.  Each source should inherit from
  # {Source} and reimplement {Source#fetch} and {Source#parse} if necessary.
  # @abstract
  class Source

    # Create a new {Source} object
    def initialize(options = {})
      @options = options
    end

    # Fetch recent entries from the source's feed
    # @return [Array<Object>] an array of {Feed::FeedEntry} objects
    def fetch
      feedjira_entries = Feedjira::Feed.fetch_and_parse(self.class.feed).entries
      feedjira_entries.map do |e|
        Feed::FeedEntry.new(
          e.url,
          self.class.title
        )
      end
    end

    # Create an {Article} object from a {Feed::FeedEntry}
    #
    # @param [Feed::FeedEntry] entry a single feed entry
    # @return [Article] an {Article} instance
    def parse(entry)
      html = open(entry.url, :allow_redirections => :safe).read
      document = Readability::Document.new(html)
      document = readability(html)
      image = document.images.empty? ? nil : document.images.first
      Article.new(
        title: document.title,
        author: document.author,
        body: document.content,
        image: image,
        source: self.class.title,
        via: entry.via,
        permalink: entry.url
        )
    end

    # Recognizes a source from an article permalink
    #
    # @param [String] permalink an article permalink
    # @return [Boolean] whether the article is recognized
    def recognize?(permalink)
      return false if !self.class.homepage
      host_and_path(permalink).start_with?(host_and_path(self.class.homepage))
    end

    class << self
      # The full title of the source
      # @note This attribute is set using a DSL
      # @example Define this attribute in child classes
      #   title "My Source"
      # @attribute [r]
      # @scope class
      # @return [String]
      attr_rw :title

      # The stylesheet to use when rendering the source
      # @note This attribute is set using a DSL
      # @example Define this attribute in child classes
      #   style "mysource"
      # @attribute [r]
      # @scope class
      # @return [String]
      attr_rw :style

      # The feed URL of the source
      # @note This attribute is set using a DSL
      # @example Define this attribute in child classes
      #   feed "http://feeds.feedburner.com/<my-source>"
      # @attribute [r]
      # @scope class
      # @return [String]
      attr_rw :feed

      # The homepage source
      # @note This attribute is set using a DSL
      # @example Define this attribute in child classes
      #   homepage "http://mysource.com"
      # @attribute [r]
      # @scope class
      # @return [String]
      attr_rw :homepage
    end

    private

    def host_and_path(uri_string)
      uri = URI(uri_string)
      uri.host + uri.path
    end

  end

  # The default source class for unrecognized articles
  class DefaultSource < Source
    def parse(entry)
      html = open(entry.url, :allow_redirections => :safe).read
      document = readability(html)
      Article.new(
        title: document[:title],
        author: document[:author],
        body: document[:content],
        source: URI(entry.url).host,
        via: entry.via,
        permalink: entry.url
        )
    end
  end

end
