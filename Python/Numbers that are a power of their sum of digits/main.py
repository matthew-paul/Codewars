def power_sumDigTerm(n):
    return sorted([i**j for j in range(2,50) for i in range(2,100) if i == sum(map(int, str(i**j)))])[n-1]