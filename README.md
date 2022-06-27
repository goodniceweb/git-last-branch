# git last branch

The plugin simplifies an access to recent branches:

## Pre-requirements

Pre-requirement for this is golang >= 1.17

Also, please ensure you have a custom folder in your home directory
where you can put the executable. In the example below it's `bin`:

```
mkdir ~/bin
echo 'export PATH=$PATH:~/bin' >> ~/.bashrc # or .zshrc or .fishrc
```

Feel free to skip this step if you already have such a folder.

## Install

### Build from source


```
git clone https://github.com/goodniceweb/git-last-branch.git
cd git-last-branch
go build
mv main ~/bin/git-last-branch
git config --global alias.lbr '!git-last-branch'
```

## Uninstall

```
rm ~/bin/git-last-branch
```

Also, please don't forget to edit your `~/.gitconfig` file and remove the `lbr` alias.
