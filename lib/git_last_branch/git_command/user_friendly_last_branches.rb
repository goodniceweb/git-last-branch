# frozen_string_literal: true

require_relative 'base_command'

module GitLastBranch
  module GitCommand
    class UserFriendlyLastBranches < BaseCommand
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
        add_git_flag("--format='%(refname)%09%(objectname:short)%09%(authordate)'")
        add_pipe_command('column -t')
        add_pipe_command("head -n #{amount}")
        to_s
      end

      private

      attr_reader :amount
    end
  end
end
