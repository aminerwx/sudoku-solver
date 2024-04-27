run:
	@go run main.go

build:
	@go build -o bin/sudoku-solver main.go

clean:
	@rm bin/sudoku-solver

all: clean build
	@./bin/sudoku-solver
