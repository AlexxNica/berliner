require "spec_helper"
require "berliner/feed"

shared_examples_for "a filter" do

  let(:input_feed) { Berliner::Feed.new([]) }

  let(:filter) do
    described_class.new
  end

  describe "#filter" do
    it "should be defined" do
      expect(filter.respond_to?(:filter)).to be true
    end

    it "should take a feed and an options hash, and return a new feed" do
      output = filter.filter(input_feed, {})
      expect(output).to be_a Berliner::Feed
      expect(output.sources).to eq(input_feed.sources)

      # Output entries should be a subset of input entries
      output.entries.each do |entry|
        expect(entry).to be_in(input_feed.entries)
      end
    end

  end
end
