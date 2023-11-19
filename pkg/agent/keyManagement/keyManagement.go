package keyManagement

import (
     "math/rand"
     log "github.com/sirupsen/logrus"
)

// Narrow to readable ascii chars
const MIN_VAL_FM int32 = 33
const MAX_VAL_FM int32 = 126

// Generate key
func Keygen(fm_matrix [][]int32,matrix [][]int32,current_update int) []int32{
    log.Info("Generating new key...")
    size := len(matrix[0])

    var i int = rand.Intn(size)
    var j int = rand.Intn(size)
    
    var key []int32 = make([]int32,size)
    for it:=0;it<size;it++{
	key[it] = fm(fm_matrix,matrix[i][it],matrix[it][j])
    }

    log.Info("Key created")
    log.Debug("Key:",key)

    return key

}

// Update Z matrix
func Update_matrix(matrix [][]int32) [][]int32 {
    
    log.Info("Updating Z matrix...")

    var size int = len(matrix[0])
     
    updated_matrix := make([][]int32,size)
    for i := range updated_matrix {
	updated_matrix[i] = make([]int32,size)
    }

    // Horizontal rotation
    for i:=0;i<size;i++{
	line := matrix[i]
	n_rotation := rand.Intn(size)
	// Perform array rotation
	line = append(line[size-n_rotation:], line[:size-n_rotation]...) 
	for j:=0;j<size;j++{
	    updated_matrix[j][i] = line[j]
	}
    }

    // Vertical rotation
    for i:=0;i<size;i++{
	column := make([]int32,size)
	
	// Gather column values
	for j:=0;j<size;j++{
	    column[j] = matrix[j][i]
	}
	n_rotation := rand.Intn(size)
	// Perform array rotation
	column = append(column[size-n_rotation:], column[:size-n_rotation]...) 
	for j:=0;j<size;j++{
	    updated_matrix[j][i] = column[j]
	}
    }
    
    log.Info("Z matrix updated")
    return updated_matrix
}

// Generate Z matrix
func Generate_matrix(fm_matrix [][]int32,master_key string) [][]int32 {
    log.Info("Generating Z matrix...")
    // Key stirng to array of chars
    var master_key_array = []int32(master_key)
    
    size := len(master_key_array)
    if size%2 !=0{
    	log.Fatal("Master key size must be even")
	// Code exits
    }
    bytes:= size/2
    
    // Split master_key, master_key = 2*bytes
    var m1 []int32 = master_key_array[:bytes]
    var m2 []int32 = master_key_array[bytes:]
    log.Debug("M1:",m1)    
    log.Debug("M2:",m2)    
    // Generate matrix with horizontal-right rotation
    var za [][]int32 = matrix_from_Hrotation_array(m1)
    log.Debug("Za:",za)
    // Generate matrix with vertical-down rotation
    var zb [][]int32 = matrix_from_Vrotation_array(m2)
    log.Debug("Zb:",zb)
    // Generate random matrix
    var zs [][]int32 = Generate_random_matrix(bytes);
    log.Debug("Zs:",zs)

    z := make([][]int32,bytes)
    for i := range z {
	z[i] = make([]int32,bytes)
    }

    for i:=0;i<bytes;i++{
	for j:=0;j<bytes;j++{
	    z[i][j] = fm(fm_matrix,fm(fm_matrix,za[i][j],zb[i][j]),za[i][j])
	}
    }
    log.Debug("Z:",z)
    log.Info("Z matrix generated") 
    return z
}

// Generate matrix from horizontaly rotatting an array
func matrix_from_Hrotation_array(root_array []int32) [][]int32 {
   
    var size int = len(root_array)
    
    matrix := make([][]int32,size)
    for i := range matrix {
	matrix[i] = make([]int32,size)
    }

    // Add to matrix
    for i:=0;i<size;i++{
	for j:=0;j<size;j++{
	    matrix[i][j] = root_array[j]
	}
	// Rotate array to the right
	root_array = append([]int32{root_array[size-1]}, root_array[:size-1]...)
    }

    return matrix
}

// Generate matrix from horizontaly rotatting an array
func matrix_from_Vrotation_array(root_array []int32) [][]int32 {
    
    var size int = len(root_array)

    matrix := make([][]int32,size)
    for i := range matrix {
	matrix[i] = make([]int32,size)
    }

    // Add to matrix
    for i:=0;i<size;i++{
	for j:=0;j<size;j++{
	    matrix[j][i] = root_array[j]
	}
	// Rotate array to the right
	root_array = append([]int32{root_array[size-1]}, root_array[:size-1]...)
    }

    return matrix
}

// Generate random 2-dimensional array with given size
func Generate_random_matrix(size int) [][]int32 {
    
    matrix := make([][]int32,size)
    for i := range matrix {
	matrix[i] = make([]int32,size)
    }


    for i:=0;i<size;i++ {
	for j:=0;j<size;j++ {
	    // Generate numbers in range MAX_VAL_FM, MIN_VAL_FM
	    matrix[i][j] =  rand.Int31n(MAX_VAL_FM-MIN_VAL_FM+1) + MIN_VAL_FM
	}
    }
    return matrix
}

func fm(fm_matrix [][]int32, line_index int32, column_index int32) int32{
    return fm_matrix[line_index-MIN_VAL_FM][column_index-MIN_VAL_FM]
}
