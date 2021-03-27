module todo/cmd/server

go 1.16

replace (
    todo/pkg/server => ../../pkg/server
    todo/pkg/handler => ../../pkg/handler
)

require (
    todo/cmd/server v0.0.1
    todo/cmd/handler v0.0.1
)