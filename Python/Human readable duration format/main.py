def format_duration(seconds):
    result = []
    minutes = seconds // 60
    hours = minutes // 60
    days = hours // 24
    years = days // 365
    result.append(seconds % 60)
    result.append(minutes % 60)
    result.append(hours % 24)
    result.append(days % 365)
    result.append(years)
    return format_string(result)

def format_string(times):
    result = []
    types = ['second', 'minute', 'hour', 'day', 'year']
    for i in range(5):
        if times[i] > 0:
            result.append(f"{times[i]} {types[i]}{'s' if times[i] >  1 else ''}")
    
    hr_time = ''
    if len(result) == 0:
        return 'now'
    elif len(result) == 1:
        return result[0]
    for i in range(len(result)-1, -1, -1):
        if i > 1:
            hr_time += result[i] + ', '
        elif i == 1:
            hr_time += result[i] + ' and '
        else:
            hr_time += result[i]
    return hr_time
        
        

    
