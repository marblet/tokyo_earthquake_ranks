module github.com/marblet/tokyo_earthquake_ranks/backend/

go 1.14

replace local.packages/apis => ./pkg/apis

require (
	github.com/ant0ine/go-json-rest v3.3.2+incompatible
	local.packages/apis v0.0.0-00010101000000-000000000000
)
