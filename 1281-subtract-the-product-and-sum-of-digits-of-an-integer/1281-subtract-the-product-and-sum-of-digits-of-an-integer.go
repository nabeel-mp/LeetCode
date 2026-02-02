func subtractProductAndSum(n int) int {
    productOfDigits := 1
    sumOfDigits := 0

    for n > 0 {
        digit := n % 10 
        productOfDigits *= digit
        sumOfDigits += digit
        n/=10
    }
    return productOfDigits - sumOfDigits
}