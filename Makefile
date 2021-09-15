.PHONY: build

# Builds the binaries.
build: clean pack
	go build -o buttress-server

# Cleans the project.
clean:
	rm -f buttress-server

# Packs the static resources.
pack:
	rm -f pkged.go
	pkger
