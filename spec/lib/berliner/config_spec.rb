require "spec_helper"

describe Berliner do
  it "should have a description" do
    expect(Berliner::DESCRIPTION).not_to be_nil
  end

  it "should have a lib path" do
    expect(Berliner::LIB_DIR).not_to be_nil
  end
end