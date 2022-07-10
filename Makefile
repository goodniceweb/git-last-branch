build:
	mkdir -p dist && go build -o dist/git-last-branch

install:
	mv dist/git-last-branch ~/bin/ && git config --global alias.lbr '!git-last-branch'