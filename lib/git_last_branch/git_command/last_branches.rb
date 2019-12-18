# frozen_string_literal: true

require_relative 'base_command'

module GitLastBranch
  module GitCommand
    class LastBranches < BaseCommand
      class << self
        # TODO: make it configurable on command line util level
        def list(amount: 10)
          new(amount: amount).list
        end
      end

      def initialize(amount: 10)
        super
        @amount = amount
      end

      def list
        add_git_flag("--format='%(refname)'")
        add_pipe_command("head -n #{amount}")
        to_s
      end

      private

      attr_reader :amount
    end
  end
end
