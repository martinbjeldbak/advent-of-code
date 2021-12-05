# Advent of Code
[![ci](https://github.com/martinbjeldbak/advent-of-code/actions/workflows/ci.yml/badge.svg)](https://github.com/martinbjeldbak/advent-of-code/actions/workflows/ci.yml)

Documenting my Advent of Code solutions while learning [golang](https://go.dev/).

This project uses the [spf13/cobra](https://pkg.go.dev/github.com/spf13/cobra) CLI package to help facilitate running of individual day solutions.

Check out the help menu for all currently implemented solutions and other commands after cloning the repository:

```sh
$ go run main.go --help
```

Note this requires golang version 1.17 installed. I use [`asdf`](https://asdf-vm.com/) to manage go versions.

For example, to run the solution for day 1, part 2 with input day1_input.txt in the [`test/`](test/) dir, run the below command.

```sh
$ go run main.go day1part2 -F test/day1_input.txt
$ # 1158
```
