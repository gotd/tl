module github.com/gotd/tl

go 1.18

require (
	github.com/sebdah/goldie/v2 v2.8.0
	github.com/stretchr/testify v1.12.1
)

require (
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sergi/go-diff v1.0.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
)

replace gopkg.in/yaml.v3 => gopkg.in/yaml.v3 v3.0.0 // CVE-2022-28948
