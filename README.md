# Go Playground (goplayground)

Fun times exploring different features, libraries, and challenges in Go. Modernized for Go 1.26.

## Sub-Projects

Below is a directory of the projects within this repository, along with brief descriptions of what they do:

*   **[fib](./fib)**: Fibonacci sequence calculators. Features recursive, parallelized recursive (using goroutines), iterative, closure-based generator, and modern Go 1.23+ `iter.Seq` iterator implementations.
*   **[fractals](./fractals)**: A Mandelbrot set visualizer using Ebitengine and a fractal tree generator using `gg`.
*   **[gameoflife](./gameoflife)**: Conway's Game of Life simulated dynamically inside the terminal using the `tcell` terminal library, modernized to range over cells using a custom `iter.Seq2` iterator.
*   **[generics](./generics)**: Exploration of Go generic type parameters. Implements an `option` type and a monadic `bind` operation.
*   **[hackerrank](./hackerrank)**: Solutions to programming challenges from HackerRank, including Euler Project #1.
*   **[http](./http)**: A command-line client demonstrating raw HTTP requests, custom redirects, and method assertions using standard library structured logging (`log/slog`).
*   **[merkle](./merkle)**: Client demonstrating cryptographic verification of serverless transparency logs by fetching checkpoints and verifying log signatures.
*   **[websockets](./websockets)**: A real-time chat application using Fiber, websockets, and HTML templates. Robustly handles errors without exiting the server.

## Getting Started

Make sure you have Go 1.26+ installed.

### Run tests
```bash
go test ./...
```
