func isPalindrome(s string) bool {
    clean := []byte{}
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c >= 'A' && c <= 'Z' {
            c += 32
        }
        if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
            clean = append(clean, c)
        }
    }
    n := len(clean)
    for left := 0; left < n/2; left++ {
        if clean[left] != clean[n-1-left] {
            return false
        }
    }
    return true
}