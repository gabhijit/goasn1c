GO_SRCS := $(shell find . -name '*.go')

all: $(GO_SRCS)
	go get && go generate && go build -v


clean:
	$(RM) -rf ./goasn1c
.PHONY: all
