module Berliner
  # Limits the number of articles from each source to a specified number
  class PerSourceLimitFilter
    attr_accessor :options

    def initialize(args = [])
      limit = args.shift.to_i

      options = {}
      options[:limit] = limit if limit > 0
      @options = default_options.merge(options)
    end

    def filter(feed)
      entries = feed.entries
                .group_by(&:via)
                .each_with_object({}) do |(source, entries_for_source),
                                          meta_entries|
                  meta_entries[source] = entries_for_source
                                         .slice(0, options[:limit])
                  meta_entries
                end
                .values
                .flatten

      feed.entries = entries
      feed
    end

    private

    def default_options
      {
        limit: 1
      }
    end
  end
end
