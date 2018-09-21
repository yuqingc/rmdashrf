ENTRY := cmd/rmdashrf/main.go
OUTDIR := bin
all: cmd/rmdashrf/main.go
	go build -o $(OUTDIR)/main.out $(ENTRY)