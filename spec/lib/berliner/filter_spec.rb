require "spec_helper"

shared_examples_for "a filter" do

  let(:input_articles) do
    articles = []
    3.times do
      articles << Berliner::Article.new(title: "Test", body: "Test")
    end
    articles
  end

  let(:filter) do
    described_class.new(input_articles)
  end

  context "initialized with 3 articles" do

    it "stores the input articles correctly" do
      expect(filter.input).to eq(input_articles)
    end

    describe "#output_articles" do
      it "returns an array of 3 or less articles" do
        expect(filter.output.size).to be <= 3
        filter.output.each do |item|
          expect(item).to be_a Berliner::Article
        end
      end
    end

  end
end

describe Berliner::Filter do
  it_behaves_like "a filter"
end
