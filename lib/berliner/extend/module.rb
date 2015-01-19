# Module additions
class Module
  # Method to implement the {Berliner::Source} DSL.  Allows definition
  # of attributes without an `=`.
  # @param [Symbol, Array<Symbol>] attrs the attributes to expose
  # @return [void]
  def attr_rw(*attrs)
    file, line, _ = caller.first.split(":")
    line = line.to_i

    attrs.each do |attr|
      module_eval <<-EOS, file, line
        def #{attr}(val=nil)
          val.nil? ? @#{attr} : @#{attr} = val
        end
      EOS
    end
  end
end
