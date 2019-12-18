# frozen_string_literal: true

require 'thor'

module GitLastBranch
  # Handle the application command line parsing
  # and the dispatch to various command objects
  #
  # @api public
  class CLI < Thor
    # Error raised by this runner
    Error = Class.new(StandardError)

    desc 'version', 'git_last_branch version'
    def version
      require_relative 'version'
      puts "v#{GitLastBranch::VERSION}"
    end
    map %w(--version -v) => :version

    desc 'show_list', 'Command description...'
    method_option :help, aliases: '-h', type: :boolean,
                         desc: 'Display usage information'
    method_option :amount, aliases: '-n', type: :numeric,
                           desc: 'Set amount of branches you want to see'
    def show_list(*)
      if options[:help]
        invoke :help, ['show_list']
      else
        require_relative 'commands/show_list'
        GitLastBranch::Commands::ShowList.new(options).execute
      end
    end
  end
end
