require "spec_helper"

shared_examples_for "a filter" do

  let(:input_articles) do
    articles = []
    3.times do
      articles << Article.new(title: "Test", body: "Test")
    end
    articles
  end

  let(:filter) do
    described_class.new(input_articles)
  end

  context "initialized with 3 articles" do

    it "stores the input articles correctly" do
      filter.input.should eq(input_articles)
    end

    describe "#output_articles" do
      it "returns an array of 3 or less articles" do
        filter.output.size.should be <= 3
        filter.output.each do |item|
          item.should be_a Article
        end
      end
    end

  end
end

describe Berliner::Filter do
  it_behaves_like "a filter"
end
