require "feedjira"
require "ruby-readability"
require "open-uri"
require "berliner/article"
require "active_support"
require "active_support/core_ext"

# The base object for a Berliner source.  Each source should inherit from
# {Source} and reimplement {Source#fetch} and {Source#parse} if necessary.
# @abstract
class Source

  # Create a new {Source} object
  def initialize(options = {})
    @options = options
  end

  # Fetch recent entries from the source's feed
  # @return [Array<Object>] an array of recent Feedjira entry object
  def fetch
    entries = Feedjira::Feed.fetch_and_parse(self.class.feed).entries

    # For now, take articles published in the last 3 days, up to a max of 3 articles
    entries.reject{ |entry| entry.published < (Time.now - 3.days) }.take(3)
  end

  # Create an {Article} object from a Feedjira entry
  # 
  # @param [Object] entry a single Feddjira entry object
  # @return [Article] an {Article} instance
  def parse(entry)
    html = open(entry.url).read
    document = Readability::Document.new(html)
    Article.new(
      title: document.title,
      author: document.author,
      content: document.content,
      source: self.class.title,
      style: self.class.style
      )
  end

  # Get all the articles for a given source
  # @note This method simply maps the entries returned by {#fetch} through {#parse}
  # @return [Array<Article>] the source's articles
  def articles
    @entries = fetch
    articles = @entries.map do |entry|
      parse(entry)
    end
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
  end

end
