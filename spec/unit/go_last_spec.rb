require 'git_last_branch/commands/go_last'

RSpec.describe GitLastBranch::Commands::GoLast do
  it "executes `go_last` command successfully" do
    output = StringIO.new
    options = {}
    command = GitLastBranch::Commands::GoLast.new(options)

    command.execute(output: output)

    expect(output.string).to eq("OK\n")
  end
end
