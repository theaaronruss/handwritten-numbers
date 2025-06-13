package main

import "fmt"

type matrix struct {
	rows int
	cols int
	data [][]float64
}

func (mat *matrix) multiply(other *matrix) matrix {
	if mat.cols != other.rows {
		panic("Matrices can not be multiplied")
	}
	result := matrix{
		mat.rows, other.cols,
		make([][]float64, mat.rows),
	}
	for row := range mat.rows {
		result.data[row] = make([]float64, other.cols)
		for col := range other.cols {
			for i := range mat.cols {
				result.data[row][col] += mat.data[row][i] * other.data[i][col]
			}
		}
	}
	return result
}

func (mat *matrix) add(other *matrix) matrix {
	if mat.rows != other.rows || mat.cols != other.cols {
		panic("Matrices can not be added")
	}
	result := matrix{
		mat.rows, mat.cols,
		make([][]float64, mat.rows),
	}
	for row := range result.rows {
		result.data[row] = make([]float64, result.cols)
		for col := range result.cols {
			result.data[row][col] = mat.data[row][col] + other.data[row][col]
		}
	}
	return result
}

func main() {
	n0 := matrix{
		1, 2,
		[][]float64{
			{3, 2,},
		},
	}
	// n1 := matrix{
	// 	1, 2,
	// 	[][]float64{
	// 		{0, 0,},
	// 	},
	// }
	w1 := matrix{
		2, 2,
		[][]float64{
			{4,1,},
			{8,10,},
		},
	}
	b1 := matrix{
		1, 2,
		[][]float64{
			{1, 4,},
		},
	}
	result := n0.multiply(&w1)
	result = result.add(&b1)
	fmt.Println(result.data)
}
