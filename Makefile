# TODO: install with prefix=$HOME/.local

packages = \
  github.com/shirou/gopsutil \
  github.com/lucasb-eyer/go-colorful \
  golang.org/x/sys/unix \
  gopkg.in/alecthomas/kingpin.v2 \
  golang.org/x/crypto/ssh/terminal \
  golang.org/x/text/width \
  github.com/rs/zerolog/log \
  github.com/mattn/go-runewidth \
  gopkg.in/libgit2/git2go.v24 
  

prefix = ./dist

prg_name = golorprompt

prg_bin = $(prefix)/bin/$(prg_name)

prg_srcs = $(wildcard cmd/golorprompt/*.go)

sys_srcs = $(wildcard sys/*.go)

seg_srcs =  $(wildcard seg/*/*.go)

segments_dir = $(prefix)/lib/golorprompt

segments_1 = $(wildcard seg/*)

blacklist = seg/git2go

segments = $(filter-out $(blacklist),$(segments_1))

segment_plugins = $(addprefix $(segments_dir)/, $(addsuffix .so, $(notdir $(segments))))

dir = XXXXXXXXXXXXXXXXXXXXXXX

version = $(shell git describe)


##############################################################################

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

get:: ## Get used go packages
	go get -v $(packages)

build:: $(prg_bin) ## Compile program and plugins

$(prg_bin): $(prg_srcs) $(sys_srcs) $(seg_srcs)
	go build -i -o $@ -ldflags="-s -w" ./cmd/golorprompt



dist:: build
	mkdir -p dist/golorprompt-$(version)
	cp $(prg) dist/golorprompt-$(version)
	cd dist && zip -r golorprompt-$(version).zip golorprompt-$(version)


# -linkshared has no effect to memory size
$(segments_dir)/%: $(sys_srcs) $(seg_srcs)
	go build -buildmode=plugin -o $@ ./seg/$(notdir $(basename $@))

debug2:: $(segment_plugins) $(prg-debug) ## debug a program
	$(prg-debug) --debug

run:: $(prg) seg-cwd seg-git seg-exitcode seg-jobs seg-user

seg-cwd: $(prg) seg-cwd-usr seg-cwd-long seg-cwd-long-single

seg-cwd-usr:
	@echo "\n\n*** $@"
	@cd /usr/local/bin && $(prg)

long_path = "/tmp/long/long/${HOME}/directory/path/with space/long/hggjkgkjhkgkgkjghkkjlkhjlkhhhlhhkhkhhklj"
seg-cwd-long:
	@echo "\n\n*** $@"
	@mkdir -p $(long_path)
	@cd $(long_path) && $(prg)

long_path_single = "/tmp/0123456789/0123456789/0123456789/0123456789/0123456789/0123456789/0123456789/0123456789/"
seg-cwd-long-single:
	@echo "\n\n*** $@"
	@mkdir -p $(long_path_single)
	@cd $(long_path_single) && $(prg)

seg-exitcode: $(prg) seg-exitcode-ok seg-exitcode-error seg-exitcode-noperm seg-exitcode-notfound \
	seg-exitcode-term

seg-exitcode-ok: $(prg)
	@echo "\n\n*** $@"
	$(prg) RC=0

seg-exitcode-error: $(prg)
	@echo "\n\n*** $@"
	$(prg) RC=1

seg-exitcode-noperm: $(prg)
	@echo "\n\n*** $@"
	$(prg) RC=126

seg-exitcode-notfound: $(prg)
	@echo "\n\n*** $@"
	$(prg) RC=127

seg-exitcode-term: $(prg)
	@echo "\n\n*** $@"
	$(prg) RC=134

seg-git:: $(prg) ahead behind

$(dir)/ahead:
	mkdir -p $@-shared
	cd $@-shared && git init --bare
	mkdir -p $@
	cd $@ && git init
	cd $@ && touch README.md
	cd $@ && git add README.md
	cd $@ && git commit -m "initial commit" || true
	cd $@ && git remote add origin ../../$@-shared || true
	cd $@ && git push --set-upstream origin master
	cd $@ && echo "second modification" >> README.md
	cd $@ && git commit -a -m "second commit" || true
	cd $@ && git push
	cd $@ && echo "third modification, not pushed" >> README.md
	cd $@ && git commit -a -m "third commit, not pushed"

ahead: $(dir)/ahead $(prg)
	@echo "\n\n*** $<"
	@cd $< && git status -s -b
	@cd $< && $(prg)

$(dir)/behind:
	mkdir -p $@-shared
	cd $@-shared && git init --bare
	#
	mkdir -p $@
	cd $@ && git init
	cd $@ && touch README.md
	cd $@ && git add README.md
	cd $@ && git commit -m "initial commit" || true
	cd $@ && git remote add origin ../../$@-shared || true
	cd $@ && git push --set-upstream origin master
	# meanwhile
	mkdir -p $@-work
	cd $@-work && git init
	cd $@-work && git remote add origin ../../$@-shared || true
	cd $@-work && git pull origin master
	cd $@-work && echo "outer modification" >> README.md
	cd $@-work && git commit -a -m "outer commit" || true
	cd $@-work && git push --set-upstream origin master
	#
	cd $@ && git fetch origin master

behind: $(dir)/behind $(prg)
	@echo "\n\n*** $<"
	@cd $< && git status -s -b
	@cd $< && $(prg)

seg-user: seg-user-root seg-user-bin seg-user-nobody

seg-user-root: $(prg)
	@echo "\n\n*** $@"
	@sudo $(prg)

seg-user-bin: $(prg)
	@echo "\n\n*** $@"
	@sudo --user=bin $(prg)

seg-user-nobody: $(prg)
	@echo "\n\n*** $@"
	@sudo --user=nobody $(prg)

seg-jobs: $(prg)
	@echo "\n\n*** $@"
	sleep 2 & $(prg)

seg-level: $(prg)
	@echo "\n\n*** $@"
	zsh -c echo $$SHLVL
