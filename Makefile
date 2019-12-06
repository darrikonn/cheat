ifeq ($(OS),Windows_NT)
EXE=cheat.exe
RM=cmd /c del
LDFLAG=
else
EXE=cheat
RM=rm
LDFLAG=-fPIC
endif

all: $(EXE) $(EXT)

$(EXE): cli/cheat.go
	go build $<

clean:
	@-$(RM) $(EXE) $(EXT)

requirements:
	go get -t

dev-requirements:
	@while read in; do \
		echo "$$in"; \
		eval "$$in"; \
	done < dev-requirements.txt

golint:
	golint ./...

vet:
	go vet ./...

errcheck:
	errcheck ./...

deadcode:
	find cli -type d | xargs deadcode

lint: golint vet errcheck deadcode

format:
	go fmt ./...
