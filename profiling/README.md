# Profiling in Go

Profiling is an essential technique for analyzing the runtime behavior and performance characteristics of your Go programs. Go provides built-in support for various types of profiling, such as CPU profiling, memory profiling, block profiling, and more. This README will guide you through using profiling in Go and running tests using the provided test files.

## Contents
- [Running Tests](#running-tests)
- [Profiling Types](#profiling-types)
- [Using Profiling](#using-profiling)
  - [CPU Profiling](#cpu-profiling)
- [Interpreting Results](#interpreting-results)
- [Running CPU Profiling Example](#running-cpu-profiling-example)

## Running Tests

To run tests for the code provided in the `main.go` and `main_test.go` files, follow these steps:

1. Make sure you have Go installed on your machine.

2. Navigate to the directory containing the `main.go` and `main_test.go` files in your terminal.

3. Run the tests using the following command:
`go test
`

This will execute the test functions defined in `main_test.go` and report any failures or successes.

## Profiling Types

Go provides several types of profiling:
- CPU Profiling: Measures CPU usage by functions.
- Memory Profiling: Analyzes memory allocations and usage.
- Block Profiling: Detects blocked goroutines.
- Goroutine Profiling: Provides information about active goroutines.

## Using Profiling

### CPU Profiling

CPU profiling helps identify where your program spends the most time. Here's how to use CPU profiling:

1. **Instrument Your Code**: Import the `net/http/pprof` package in your code to expose profiling endpoints. Example: `_ "net/http/pprof"`.

2. **Start Profiling Server**: In your `main` function, start an HTTP server to expose profiling endpoints.

3. **Run Your Program**: Execute your Go program.

4. **Access Profiling Data**: Open a web browser and navigate to `http://localhost:6060/debug/pprof/`. Click "CPU Profile" to generate the profile data.

### Interpreting Results

Profiling data helps you identify bottlenecks and areas for optimization in your code. Use tools like the pprof web UI and CLI to analyze the data.

## Running CPU Profiling Example

The provided test files include an example of running CPU profiling:
- The `TestMain` function starts the HTTP server for profiling endpoints.
- The `TestHTTPServer` function tests the profiling endpoint's response.
- The `TestCompute` function tests the `compute` function.

To run the CPU profiling example:

1. Navigate to the directory containing the `main.go` and `main_test.go` files.

2. Run the tests using the following command:

`go test -cpuprofile cpu.prof`

3. After running the tests, you'll have a `cpu.prof` file. Use the `go tool pprof` command to analyze the profile data:

`go tool pprof cpu.prof`

The pprof CLI will open, allowing you to explore the CPU profiling data.
