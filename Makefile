build:
		go build -o bin/nmath main.go

run:
		go run main.go

test:
		cd nmath && go test

coverage:
		cd nmath && go test -coverprofile=coverage.out

clean:
		rm -fr bin nmath/coverage.out
