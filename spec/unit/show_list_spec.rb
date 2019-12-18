require 'git_last_branch/commands/show_list'

RSpec.describe GitLastBranch::Commands::ShowList do
  it "executes `show_list` command successfully" do
    output = StringIO.new
    options = {}
    command = GitLastBranch::Commands::ShowList.new(options)

    command.execute(output: output)

    expect(output.string).to eq("OK\n")
  end
end
