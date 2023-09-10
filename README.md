# go-lang-learn

My attempt at learning go

## Setup

<https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code#4-update-the-go-tools>

## Learning/general
<https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection>

## Modules

go mod init mymodulequalifiedname
go work init
inside go.work:
    use .

## Test run examples

go test utils/*
go test utils/* -v
go test utils/*-v -cover
go test utils/* -v -bench

## Run all tests

go test -v ./...

## Various example commands

go test -run TestParallelExecutor -v
go run .
