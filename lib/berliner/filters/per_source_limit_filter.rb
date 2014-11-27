module Berliner

  # Limits the number of articles from each source to a specified number
  class PerSourceLimitFilter

    def filter(articles, options={})
      options = default_options.merge(options)

      articles.
        group_by { |art| art.source }.
        inject({}) do |articles, (source, articles_for_source)|
          articles[source] = articles_for_source.slice(0, options[:limit])
          articles
        end.
        values.
        flatten!
    end

    private

    def default_options
      {
        limit: 1
      }
    end

  end
end
