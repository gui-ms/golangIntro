## Tips for making it work on windows:

- To run the command `go build`, you need to have a .mod file. This can be done by running the command `go mod init moduleName`;
  
- I also had to specify what kind of file I was outputting when running the build command. If I didn't do that the file generated wouldn't work as intended. Example: `go build -o name.exe`;
  
- To run benchstat you may need to install the package first. This can be done by running the command `go install golang.org/x/perf/cmd/benchstat@latest`;
  
- You can run a benchmark by typing the command `go test -bench=. -benchmem -count 10 > 1.bench`. This will create a file called 1.bench which will run a function a number of times defined by count. I also had to specify
which function I wanted in the `-bench` flag (in my case it was BenchmarkGenerateLargeString), or it just wouldn't run;

- To use Benchstat just run `benchstat benchFile1.bench benchFile2.bench`.
