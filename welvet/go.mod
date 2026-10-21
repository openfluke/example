module github.com/openfluke/example/welvet

go 1.22.5

require github.com/openfluke/welvet v1.1.7

require github.com/openfluke/webgpu v1.0.5 // indirect

// Until v1.1.7 is on the proxy, use the sibling engine checkout:
replace github.com/openfluke/welvet => ../../welvet
