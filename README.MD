## Using software

### With docker

Prerequisites: docker

```
docker build -t matrix_test_task .
docker run -i matrix_test_task 
```

### Without docker

Prerequisites: go version >= 1.17 installed

Run program via:

```
go run ./cmd/matrix_solver
```

### Example usage:

```2
3 4
1000
0010
0000
0 1 1 2
1 1 0 1
2 2 1 2

3 4
0000
0000
0001
5 4 3 2
4 3 2 1
3 2 1 0
```
### Note
I do understand that whole thing could've been packed into main.go/main_test.go files without all that boilerplate :)