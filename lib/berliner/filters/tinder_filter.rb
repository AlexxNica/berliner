module Berliner
  # Limits the number of articles from each source to a specified number
  class TinderFilter
    attr_accessor :options

    def initialize(args = [])
    end

    def filter(feed)
      entries = []

      feed.entries.each do |entry|
        puts ""
        puts "#{entry.via} - #{entry.title}"
        puts "Type y to swipe right."

        entries << entry if gets.chomp == "y"
      end

      feed.entries = entries
      feed
    end

    private

    def default_options
    end
  end
end
