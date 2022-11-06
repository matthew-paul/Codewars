def count_change(money, coins, index=0):
    if money == 0:
        return 1
    
    combinations = 0
    for i in range(index, len(coins)):
        if money - coins[i] >= 0:
            combinations += count_change(money - coins[i], coins, i)
    
    return combinations