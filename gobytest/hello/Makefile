dist_dir=.dist

build:
	rm -rf $(dist_dir)
	mkdir $(dist_dir)
	date > $(dist_dir)/build.info.txt
	go build -o $(dist_dir) -ldflags " -X 'main.version=$$(cat version)'"

test:
	go test ./...

test-bench:
	go test -bench=. ./...

lint:
	go vet ./...
	errcheck ./...  
