require "berliner/article"

module Berliner
  # Filters are composable objects that get passed an array of {Article} objects
  # upon initialization. They define an `output` method which returns an array
  # of {Article} objects that is a subset of the original input array, filtered
  # and ordered in some way.

  # This filter serves as an example "no-op" filter, that just returns the
  # input array as the output. It exists to be extended by filters with
  # actual functionality.
  class Filter
    attr_accessor :input

    def initialize(articles)
      @input = articles
    end

    def output
      input
    end

  end
end
