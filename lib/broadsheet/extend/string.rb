class String
  def format_heredoc
    gsub(/^[\s\t]*/, '').gsub(/[\s\t]*\n/, ' ').strip
  end
end
