def remove_smallest(numbers):
    if len(numbers) <= 1:
        return []
    smallest = numbers.index(min(*numbers))
    return [numbers[i] for i in range(len(numbers)) if i != smallest]
    
