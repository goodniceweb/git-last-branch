# frozen_string_literal: true

module GitLastBranch
  module GitCommand
    class BaseCommand
      GIT_COMMAND = 'git for-each-ref'
      GIT_OPTION = 'refs/heads'

      def initialize(*)
        @git_flags = ["--sort='-authordate'"]
        @pipe_commands = ['', "sed -e 's-refs/heads/--'"]
      end

      def add_git_flag(option)
        @git_flags.push(option)
      end

      def add_pipe_command(command)
        @pipe_commands.push(command)
      end

      def to_s
        "#{GIT_COMMAND}"\
          " #{git_flags.join(' ')}"\
          " #{GIT_OPTION}"\
          " #{pipe_commands.join(' | ')}" 
      end

      private

      attr_reader :git_command, :git_flags, :pipe_commands
    end
  end
end
