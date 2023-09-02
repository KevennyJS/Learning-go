module example.com/hello

go 1.21.0

require golang.org/x/example/hello v0.1.0 // indirect

replace golang.org/x/example/hello => ../example/hello/

//replace dependence
//build replaced version:
//go get golang.org/x/example/hello@v0.1.0