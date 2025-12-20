func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    visited := make([]bool, n)
    count := 0

    for i :=0; i<n; i++{
        if !visited[i]{
            dfsprovince(i, isConnected, visited)
            count++
        }
    }
    return count
}

func dfsprovince(city int, isConnected [][]int, visited[]bool ){
    visited[city] =true

for neigh :=0; neigh <len(isConnected); neigh++ {
    if isConnected[city][neigh]==1 && !visited[neigh] {
        dfsprovince(neigh, isConnected ,visited )   
    }
}
    
}