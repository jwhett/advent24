.PHONY: new today
.DEFAULT: fzf

BRANCH := $(shell date "+%Y%b%d")

new:
	git sw -c $(BRANCH)

today:
	git sw $(BRANCH)
