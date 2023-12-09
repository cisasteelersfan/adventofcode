#!/bin/bash

: ${1?"Usage: $0 day"}

mkdir -p "2023/day$1"
touch "2023/day$1/sol.go"
touch "2023/day$1/small.txt"
touch "2023/day$1/input.txt"

echo """package main

import \"os\"

func main() {
	dat, _ := os.ReadFile(\"2023/day$1/small.txt\")
}""" >> "2023/day$1/sol.go"
