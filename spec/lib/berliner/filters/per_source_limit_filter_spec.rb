require "spec_helper"
require "berliner/feed"
require "berliner/filters/per_source_limit_filter"

describe Berliner::PerSourceLimitFilter do
  it_behaves_like "a filter"

  let(:input_feed) do
    feed = Berliner::Feed.new([])
    feed.entries = [
      Berliner::Feed::FeedEntry.new(
        "",
        "Source 1",
        ""
        ),
      Berliner::Feed::FeedEntry.new(
        "",
        "Source 1",
        ""
        ),
      Berliner::Feed::FeedEntry.new(
        "",
        "Source 1",
        ""
        ),
      Berliner::Feed::FeedEntry.new(
        "",
        "Source 2",
        ""
        )
    ]
    feed
  end

  describe "#filter" do
    it "should filter number of articles in each source to specified limit" do
      filter = Berliner::PerSourceLimitFilter.new(["2"])
      output = filter.filter(input_feed)
      expect(output.entries.select { |entry| entry.via == "Source 1" }.size).to eq(2)
      expect(output.entries.select { |entry| entry.via == "Source 2" }.size).to eq(1)
    end

    it "should default to limiting to 1 article per source if no limit is" \
       "specified" do
      filter = Berliner::PerSourceLimitFilter.new
      output = filter.filter(input_feed)
      expect(output.entries.select { |entry| entry.via == "Source 1" }.size).to eq(1)
      expect(output.entries.select { |entry| entry.via == "Source 2" }.size).to eq(1)
    end
  end
end
