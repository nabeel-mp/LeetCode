func longestPalindrome(s string) string {
    if len(s) <= 1 {
        return s
    }

    start, maxlength := 0,1
    for i :=0 ; i<=len(s); i++ {
        len1 := expandAroundCenter(s,i,i)
        len2 := expandAroundCenter(s,i,i+1)

        maxlen := len1
        if len2 > maxlen {
            maxlen =len2
        }

        if maxlen > maxlength {
            maxlength =maxlen
            start =i -(maxlen -1) /2
        }
    }
    return s[start : start+maxlength]
}

func expandAroundCenter(s string, left, right int) int{
   for left >= 0 && right < len(s) && s[left] == s[right] {
    left--
    right++
   }
   return right - left -1
}
