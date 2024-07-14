module GolangStudy/init

go 1.22.5

replace GolangStudy/init/InitLib1 => ./InitLib1

replace GolangStudy/init/InitLib2 => ./InitLib2

require (
	GolangStudy/init/InitLib1 v0.0.0-00010101000000-000000000000
	GolangStudy/init/InitLib2 v0.0.0-00010101000000-000000000000
)
