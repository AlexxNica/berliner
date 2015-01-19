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
    attr_accessor :creds, :authenticated

    # Create a new {Source} object
    # @note See {SourceManager#initialize} for the difference between
    #   "credentials" and "creds"
    def initialize(creds=nil)
      @creds = creds
      @authenticated = creds ? auth : false
    end

    # Perform source-related authorization based on available credentials
    # @return [Boolean] whether the authorization was successful or not
    def auth
      return false
    end

    # Fetch recent entries from the source's feed
    # @return [Array<String>] an array of article permalinks
    def fetch
      feedjira_entries = Feedjira::Feed.fetch_and_parse(self.class.feed).entries
      feedjira_entries.map{ |e| e.url }
    end

    # Create an {Article} object from a {Feed::FeedEntry}
    #
    # @param [Feed::FeedEntry] entry a single feed entry
    # @return [Article] an {Article} instance
    def parse(entry)
      html = open(entry.url, :allow_redirections => :safe).read
      document = Readability::Document.new(html)
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
  end

  # The default source class for unrecognized articles
  class DefaultSource < Source
    # The only difference from above is the article source
    def parse(entry)
      html = open(entry.url, :allow_redirections => :safe).read
      document = Readability::Document.new(html)
      image = document.images.empty? ? nil : document.images.first
      Article.new(
        title: document.title,
        author: document.author,
        body: document.content,
        image: image,
        source: URI(entry.url).host,
        via: entry.via,
        permalink: entry.url
        )
    end
  end

end
