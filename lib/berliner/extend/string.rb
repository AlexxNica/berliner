# String additions
class String
  # Strip leading indentation from heredoc lines
  # @return [self] heredoc stipped of leading intendation
  def format_heredoc
    gsub(/^[\s\t]*/, "").gsub(/[\s\t]*\n/, " ").strip
  end

  # Turn a query string into a regex.  If string is surrounded by
  # slashes, it is interpreted to be a regex already.
  # @return [Regexp] query regex
  def query_regex
    case self
      when %r{^/(.*)/$} then Regexp.new($1)
      else /.*#{Regexp.escape(self)}.*/i
    end
  end

  # Turn a filename into a slug
  # @return [String] a dashed slug
  def slugify
    self.gsub(/_/, "-")
  end

  # Turn a slug into an filename
  # @return [Regexp] an underscored filename
  def deslugify
    self.gsub(/-/, "_")
  end
end
