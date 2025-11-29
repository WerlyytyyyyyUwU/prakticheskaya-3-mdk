package main

import "fmt"

func read(n int) [][]float64 {
	m := make([][]float64, n)
	for i := 0; i < n; i++ {
		m[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&m[i][j])
		}
	}
	return m
}

func printM(m [][]float64) {
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%8.3f", m[i][j])
		}
		fmt.Println()
	}
}

func add(a, b [][]float64) [][]float64 {
	n := len(a)
	r := make([][]float64, n)
	for i := 0; i < n; i++ {
		r[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			r[i][j] = a[i][j] + b[i][j]
		}
	}
	return r
}

func scalar(a [][]float64, k float64) [][]float64 {
	n := len(a)
	r := make([][]float64, n)
	for i := 0; i < n; i++ {
		r[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			r[i][j] = a[i][j] * k
		}
	}
	return r
}

func mul(a, b [][]float64) [][]float64 {
	n := len(a)
	r := make([][]float64, n)
	for i := 0; i < n; i++ {
		r[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			s := 0.0
			for t := 0; t < n; t++ {
				s += a[i][t] * b[t][j]
			}
			r[i][j] = s
		}
	}
	return r
}

func main() {
	var op, n int
	fmt.Println("Выберите операцию:")
	fmt.Println("1 = сложение")
	fmt.Println("2 = умножение матрицы на число")
	fmt.Println("3 = умножение матриц")
	fmt.Scan(&op)

	fmt.Println("Размер (2 или 3):")
	fmt.Scan(&n)
	if n != 2 && n != 3 {
		fmt.Println("Неподдерживаемый размер")
		return
	}

	switch op {
	case 1:
		fmt.Println("Матрица A:")
		A := read(n)
		fmt.Println("Матрица B:")
		B := read(n)
		printM(add(A, B))
	case 2:
		fmt.Println("Матрица A:")
		A := read(n)
		fmt.Println("Число k:")
		var k float64
		fmt.Scan(&k)
		printM(scalar(A, k))
	case 3:
		fmt.Println("Матрица A:")
		A := read(n)
		fmt.Println("Матрица B:")
		B := read(n)
		printM(mul(A, B))
	default:
		fmt.Println("Неизвестная операция")
	}
}
