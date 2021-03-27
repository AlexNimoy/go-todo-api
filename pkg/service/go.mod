module service

go 1.16

replace (
	todo/pkg/repository => ../repository
)

require (
	todo/pkg/repository v0.0.1
)