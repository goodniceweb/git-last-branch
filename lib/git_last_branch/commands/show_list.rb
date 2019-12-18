# frozen_string_literal: true

require_relative '../command'
require_relative '../git_command/last_branches'
require_relative '../git_command/user_friendly_last_branches'
require 'pry'

module GitLastBranch
  module Commands
    class ShowList < GitLastBranch::Command
      GitCommandError = Class.new(StandardError)

      extend Forwardable

      def_delegators :cmd, :run
      def_delegators :@prompt, :select

      def initialize(options)
        # TODO: make it configurable on command line util level
        @amount = options.fetch(:amount, 10)
        # TODO: make it different depends on environment
        @cmd = command(printer: :null)
        @prompt = add_vim_like_keybindings
      end

      def execute(input: $stdin, output: $stdout)
        result = select('Select branch:', per_page: amount) do |menu|
          menu.choices choices_mapping
        end
        run "git checkout #{result}"
        output.puts "Checkout to #{result} branch"
      end

      private

      attr_reader :cmd, :amount

      def choices_mapping
        branches = git_branches(GitLastBranch::GitCommand::LastBranches.list(amount: amount))
        user_branches = git_branches(GitLastBranch::GitCommand::UserFriendlyLastBranches.list(amount: amount))
        Hash[user_branches.zip(branches)]
      end

      def git_branches(script)
        branches, error = run(script)
        raise GitCommandError, error unless error.strip.empty?

        branches.split("\n")
      end

      def add_vim_like_keybindings
        prompt.tap do |prompt_config|
          prompt_config.on(:keypress) do |event|
            if event.value == 'j'
              prompt_config.trigger(:keydown)
            end

            if event.value == 'k'
              prompt_config.trigger(:keyup)
            end
          end
        end
      end
    end
  end
end
