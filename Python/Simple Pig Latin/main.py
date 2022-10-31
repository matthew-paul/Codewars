def pig_it(text):
    return ' '.join([x[1:]+x[0]+('ay' if x[0] not in ['.', ',', '!', '?'] else '') for x in text.split(' ')])