package agent

import (
    "fmt"
    "math/rand"
)

const BYTES = 7
// Narrow to readable ascii chars
const MIN_VAL_FM int = 33
const MAX_VAL_FM int = 126

func Generate_matrix(fm_matrix [][]int,master_key string) [][]int{
    // Key stirng to array of chars
    master_key_array = []rune(master_key)
    
    // Split master_key, master_key = 2*bytes
    m1 := master_key_array[:BYTES]
    m2 := master_key_array[BYTES:]
    
    // Matrix with horizontal-right rotation
    za := matrix_horizontal_rotation(m1,1)
    // Matrix with vertical-down rotation
    zb := matrix_vertical_rotation(m2,1)

    zs := generate_random_matrix();


}

// Generate matrix from horizontaly rotatting an array
func matrix_from_Hrotation_array(root_array []int,n_rotations int) [][]int {
   
    matrix := make([][]int,size)
    for i := range matrix {
        matrix[i] = make([]int,size)
    }

    // Add to matrix
    for i:=0;i<size;i++{
	for j:=0,j<size;j++{
	    matrix[i][j] = root_array[j]
	}
	// Rotate array to the right

    }

    return matrix
}


// Generate random 2-dimensional array with given size
func generate_random_matrix(size int) [][]int {
    
    matrix := make([][]int, size)
    for i := range matrix {
        matrix[i] = make([]int, size)
    }

    for i:=0;i<size;i++ {
	for j:=0;j<size;j++ {
	    // Generate numbers in range MAX_VAL_FM, MIN_VAL_FM
	    matrix[i][j] =  rand.Intn(MAX_VAL_FM-MIN_VAL_FM+1) + MIN_VAL_FM
	}
    }
    return matrix
}


