#!/usr/bin/make -f

dir=samplegits
prompt=$(PWD)/ingoline

all: empty untracked unstaged staged stash $(dir)/shared ahead behind merge detached

$(dir)/empty:
	mkdir -p $@
	cd $@ && git init

empty: $(dir)/empty
	@echo "\n*** $<"
	@cd $< && git status -s
	@cd $< && $(prompt)


$(dir)/untracked:
	mkdir -p $@
	cd $@ && git init
	cd $@ && touch README.md

untracked: $(dir)/untracked
	@echo "\n\n*** $<"
	@cd $< && git status -s
	@cd $< && $(prompt)

$(dir)/staged:
	mkdir -p $@
	cd $@ && git init
	cd $@ && touch README.md
	cd $@ && git add README.md

staged: $(dir)/staged
	@echo "\n\n*** $<"
	@cd $< && git status -s
	@cd $< && $(prompt)

$(dir)/unstaged:
	mkdir -p $@
	cd $@ && git init
	cd $@ && touch README.md
	cd $@ && git add README.md
	cd $@ && git commit -m "test"
	cd $@ && echo "thing" >> README.md

unstaged: $(dir)/unstaged
	@echo "\n\n*** $<"
	@cd $< && git status -s
	@cd $< && $(prompt)

$(dir)/stash:
	mkdir -p $@
	cd $@ && git init
	cd $@ && touch README.md
	cd $@ && git add README.md
	cd $@ && git commit -m "test"
	cd $@ && git stash
	cd $@ && echo "thing" >> README.md
	cd $@ && git stash

stash: $(dir)/stash
	@echo "\n\n*** $<"
	@cd $< && git status -s
	@cd $< && $(prompt)


$(dir)/shared:
	mkdir -p $@
	cd $@ && git init --bare


$(dir)/ahead:
	mkdir -p $@
	cd $@ && git init
	cd $@ && touch README.md
	cd $@ && git add README.md
	cd $@ && git commit -m "test" || true
	cd $@ && git remote add origin ../shared || true
	cd $@ && git push --set-upstream origin master
	cd $@ && echo "thing" >> README.md
	cd $@ && git commit -a -m "test" || true
	cd $@ && git push
	cd $@ && echo "thing 2" >> README.md
	cd $@ && git commit -a -m "test 2"

ahead: $(dir)/ahead
	@echo "\n\n*** $<"
	@cd $< && git status -s
	@cd $< && $(prompt)


$(dir)/behind:
	mkdir -p $@
	cd $@ && git init
	cd $@ && touch README.md
	cd $@ && git add README.md
	cd $@ && git commit -m "test" || true
	cd $@ && git remote add origin ../shared || true
	cd $@ && git fetch origin master
	cd $@ && git branch --set-upstream-to=origin/master master

behind: $(dir)/behind
	@echo "\n\n*** $<"
	@cd $< && git status -s
	@cd $< && $(prompt)


$(dir)/merge:
	mkdir -p $@
	cd $@ && git init
	cd $@ && echo testing > README.md
	cd $@ && git add README.md
	cd $@ && git commit -m "test 1"
	cd $@ && git checkout -b second
	cd $@ && echo tetsing > README.md
	cd $@ && git commit -am "test 2"
	cd $@ && git checkout master
	cd $@ && echo tetssng > README.md
	cd $@ && git commit -am "test 3"
	-cd $@ && git merge second

merge: $(dir)/merge
	@echo "\n\n*** $<"
	@cd $< && git status -s
	@cd $< && $(prompt)

$(dir)/detached:
	mkdir -p $@
	cd $@ && git init
	cd $@ && echo testing > README.md
	cd $@ && git add README.md
	cd $@ && git commit -m "test"
	cd $@ && git checkout --detach

detached: $(dir)/detached
	@echo "\n\n*** $<"
	@cd $< && git status -s
	@cd $< && $(prompt)

