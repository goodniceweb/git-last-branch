RSpec.describe "`git_last_branch show_list` command", type: :cli do
  it "executes `git_last_branch help show_list` command successfully" do
    output = `git_last_branch help show_list`
    expected_output = <<-OUT
Usage:
  git_last_branch show_list

Options:
  -h, [--help], [--no-help]  # Display usage information

Command description...
    OUT

    expect(output).to eq(expected_output)
  end
end
