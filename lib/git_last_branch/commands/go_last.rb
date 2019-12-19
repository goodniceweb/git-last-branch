# frozen_string_literal: true

require_relative '../command'
require_relative '../git_command/last_branches'
require 'pry'

module GitLastBranch
  module Commands
    class GoLast < GitLastBranch::Command
      GitCommandError = Class.new(StandardError)

      def_delegators :cmd, :run

      def initialize(options)
        @options = options
        # TODO: make it different depends on environment
        @cmd = command(printer: :null)
      end

      def execute(input: $stdin, output: $stdout)
        branches = git_branches(GitLastBranch::GitCommand::LastBranches.list(amount: 2))
        current_branch = git_branches('git rev-parse --abbrev-ref HEAD').first
        destination_branch = branches.find { |branch| branch != current_branch }
        run "git checkout #{destination_branch}"
        output.puts "Switch to #{destination_branch}"
      end

      private

      attr_reader :cmd

      def git_branches(script)
        branches, error = run(script)
        raise GitCommandError, error unless error.strip.empty?

        branches.split("\n")
      end
    end
  end
end
