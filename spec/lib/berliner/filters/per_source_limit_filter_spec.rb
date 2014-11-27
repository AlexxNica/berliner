require "spec_helper"
require "berliner/filters/per_source_limit_filter"

describe Berliner::PerSourceLimitFilter do
  it_behaves_like "a filter"

  let(:input_articles) do
    articles = []
    3.times do |source_number|
      3.times do |article_number|
        articles << Berliner::Article.new(
          title: "Test #{article_number}",
          body: "Test",
          source: "Source #{source_number}"
        )
      end
    end

    articles
  end

  let(:filter) do
    Berliner::PerSourceLimitFilter.new
  end

  describe "#filter" do
    it "should filter number of articles in each source to specified limit" do
      output = filter.filter(input_articles, {limit: 2})
      expect(output.size).to eq(6)
      expect(output.select{ |art| art.source == "Source 1" }.size).to eq(2)
    end

    it "should default to limiting to 1 article per source if no limit is" \
       "specified" do
      output = filter.filter(input_articles)
      expect(output.size).to eq(3)
      expect(output.select{ |art| art.source == "Source 1" }.size).to eq(1)
    end
  end
end
