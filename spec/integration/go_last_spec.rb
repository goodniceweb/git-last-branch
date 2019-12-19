RSpec.describe "`git_last_branch go_last` command", type: :cli do
  it "executes `git_last_branch help go_last` command successfully" do
    output = `git_last_branch help go_last`
    expected_output = <<-OUT
Usage:
  git_last_branch go_last

Options:
  -h, [--help], [--no-help]  # Display usage information

Command description...
    OUT

    expect(output).to eq(expected_output)
  end
end
