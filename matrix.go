package matrix

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// Matrix is a 2D slice of float64 representing a matrix.
type Matrix [][]float64

// New creates a new matrix with the specified number of rows and columns,
// filled with random values between 0 and 1.
func New(rows, cols int) Matrix {
    rand.Seed(time.Now().UnixNano())
    m := make(Matrix, rows)
    for i := range m {
        m[i] = make([]float64, cols)
        for j := range m[i] {
            m[i][j] = rand.Float64()
        }
    }
    return m
}

// Add returns the sum of two matrices of the same size using multiple goroutines.
func Add(a, b Matrix) Matrix {
    if len(a) != len(b) || len(a[0]) != len(b[0]) {
    	fmt.Println("matrix sizes do not match")
	return
    }
    c := make(Matrix, len(a))
    var wg sync.WaitGroup
    for i := range a {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            c[i] = make([]float64, len(a[i]))
            for j := range a[i] {
                c[i][j] = a[i][j] + b[i][j]
            }
        }(i)
    }
    wg.Wait()
    return c
}

// Multiply returns the product of two matrices where the number of columns in the first matrix
// matches the number of rows in the second matrix using multiple goroutines.
func Multiply(a, b Matrix) Matrix {
    if len(a[0]) != len(b) {
        fmt.Println("matrix sizes do not match")
	return
    }
    c := make(Matrix, len(a))
    var wg sync.WaitGroup
    for i := range a {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            c[i] = make([]float64, len(b[0]))
            for j := range b[0] {
                for k := range b {
                    c[i][j] += a[i][k] * b[k][j]
                }
            }
        }(i)
    }
    wg.Wait()
    return c
}

//Transpose return the transposed matrix 
func Transpose(a Matrix) Matrix {
	rows := len(a)
	cols := len(a[0])

	transposed := make(Matrix, cols)
	for i := range transposed {
		transposed[i] = make([]float64, rows)
	}

	var wg sync.WaitGroup
	for i := 0; i < cols; i++ {
		wg.Add(1)
		go func(col int) {
			defer wg.Done()
			for row := 0; row < rows; row++ {
				transposed[col][row] = a[row][col]
			}
		}(i)
	}
	wg.Wait()

	return transposed
}

// TimeAdd returns the sum of two matrices and the time taken to compute it using multiple goroutines.
func TimeAdd(a, b Matrix) (Matrix, time.Duration) {
    start := time.Now()
    c := Add(a, b)
    elapsed := time.Since(start)
    return c, elapsed
}

// TimeMultiply returns the product of two matrices and the time taken to compute it using multiple goroutines.
func TimeMultiply(a, b Matrix) (Matrix, time.Duration) {
    start := time.Now()
    c := Multiply(a, b)
    elapsed := time.Since(start)
    return c, elapsed
}

// TimeTranspose returns the new transposed matrix and the time taken to compute it. 
func TimeTranspose(a Matrix) (Matrix, time.Duration) {
    start := time.Now()
    transposed := Transpose(a)
    elapsed := time.Since(start)
    return transposed, elapsed
}
