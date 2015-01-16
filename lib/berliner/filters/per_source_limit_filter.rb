module Berliner

  # Limits the number of articles from each source to a specified number
  class PerSourceLimitFilter

    def filter(feed, options={})
      options = default_options.merge(options)

      entries = feed.entries.
        group_by { |entry| entry.via }.
        inject({}) do |entries, (source, entries_for_source)|
          entries[source] = entries_for_source.slice(0, options[:limit])
          entries
        end.
        values.
        flatten

      feed.entries = entries
      feed
    end

    private

    def default_options
      {
        limit: 2
      }
    end

  end
end
