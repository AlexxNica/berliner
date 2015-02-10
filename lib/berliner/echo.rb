require "rainbow"

module Berliner
  class Echo

    # Verbosity enums
    OFF = 0
    NORMAL = 1
    VERBOSE = 2

    # Color enums
    WHITE = :white
    BLUE = :blue
    GREEN = :green
    YELLOW = :yellow
    RED = :red

    class << self

      attr_writer :verbosity

      def verbosity
        @verbosity || Echo::OFF
      end

      def debug(message)
        print(message, Echo::WHITE, Echo::VERBOSE)
      end

      def info(message)
        print(message, Echo::BLUE)
      end

      def success(message)
        print(message, Echo::GREEN)
      end

      def warn(message)
        print(message, Echo::YELLOW)
      end

      def error(message)
        print(message, Echo::RED)
      end

      # Add a source to the profile if the source is valid
      # @param [String] message a message to print to stdout
      # @param [Number] color the color of the message
      # @param [Number] v the verbosity level of the message
      # @return [void]
      def print(message, color=Echo::WHITE, verbosity=Echo::NORMAL)
        puts Rainbow(message).color(color) unless self.verbosity < verbosity
      end

    end
  end
end