class String
  def format_heredoc
    gsub(/^[\s\t]*/, "").gsub(/[\s\t]*\n/, " ").strip
  end

  def query_regex
    case self
      when %r{^/(.*)/$} then Regexp.new($1)
      else /.*#{Regexp.escape(self)}.*/i
    end
  end
end
