# git last branch

The plugin simplifies an access to recent branches:

![git-last-branch-demo](https://user-images.githubusercontent.com/5359832/176028834-bd2d74dd-4de5-486b-b399-f6bf88969a78.gif)

## Pre-requirements

Pre-requirement for this is golang >= 1.17

Also, please ensure you have a custom folder in your home directory
where you can put the executable. In the example below it's `bin`:

```
mkdir -p ~/bin
```

### bash or zsh

```
echo 'export PATH=$PATH:~/bin' >> ~/.bashrc # or .zshrc
```

Feel free to skip this step if you already have such a folder.

### fish

```
fish_add_path ~/bin
```

## Install

### Build from source

```
git clone https://github.com/goodniceweb/git-last-branch.git
cd git-last-branch
make && make install
```

## Uninstall

```
rm ~/bin/git-last-branch
```

Also, please don't forget to edit your `~/.gitconfig` file and
remove the `lbr` alias.
