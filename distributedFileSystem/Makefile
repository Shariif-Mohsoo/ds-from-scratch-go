# 'build' is a task (also called a target) that compiles our Go code
# The '@' before the command hides the command itself in the terminal output
# 'go build -o bin/fs' → build the Go project and put the compiled binary into the 'bin' folder, naming it 'fs'
build:
	@go build -o bin/fs

# 'run' is a task that depends on 'build'
# That means before running the app, it will build it first
# After building, it runs the compiled binary './bin/fs'
run: build
	@./bin/fs

# 'test' is a task to run all tests in the project
# 'go test ./...' → find and run all tests in every folder
# '-v' → verbose mode (show details of each test being run)
test:
	@go test ./... 
