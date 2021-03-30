module handler

go 1.16

replace (
    todo/pkg/service => ../service
    todo/pkg/repository => ../repository
    todo/pkg/model => ../model
    "todo/docs" => ../docs
)

require (
    todo/pkg/service v0.0.1
    todo/pkg/repository v0.0.1
    todo/pkg/model v0.0.1
)