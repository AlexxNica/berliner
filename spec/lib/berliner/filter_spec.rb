require "spec_helper"

shared_examples_for "a filter" do

  let(:input_articles) do
    articles = []
    3.times do
      articles << Berliner::Article.new(
        title: "Test",
        body: "Test",
        source: "Test Source"
      )
    end
    articles
  end

  let(:filter) do
    described_class.new
  end

  describe "#filter" do
    it "should be defined" do
      expect(filter.respond_to?(:filter)).to be true
    end

    it "should take an array of articles and an options hash, and return an" \
       "array of articles" do
      output = filter.filter(input_articles, {})
      expect(output).to be_an Array
      output.each do |article|
        expect(article).to be_a Berliner::Article
      end
    end
  end
end
