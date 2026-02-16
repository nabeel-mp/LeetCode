func countKeyChanges(s string) int {
    if len(s) < 2 {
        return 0
    }

    changes := 0
    last := strings.ToLower(string(s[0]))

    for i := 1 ; i <len(s); i++ {
        current := strings.ToLower(string(s[i]))

        if current != last {
            changes++

            last = current
        }
    }
    return changes
}