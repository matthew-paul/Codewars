def is_interesting(number, awesome_phrases):
    if interesting(number, awesome_phrases):
        return 2
    elif interesting(number+1, awesome_phrases) or interesting(number+2, awesome_phrases):
        return 1
    else:
        return 0
    

def interesting(number, awesome_phrases):
    number_str = str(number)
    if len(number_str) < 3:
        return False
    if number in awesome_phrases:
        return True
    if all(x == '0' for x in number_str[1:]) and number_str[0] != '0':
        return True
    if all(x == number_str[0] for x in number_str):
        return True
    if all(int(number_str[i]) == int(number_str[i-1])+1 or (number_str[i] == '0' and number_str[i-1] == '9')
           for i in range(1, len(number_str))):
        return True
    if all(int(number_str[i]) == int(number_str[i-1])-1 for i in range(1, len(number_str))):
        return True
    if all(number_str[i] == number_str[len(number_str)-1-i] for i in range(len(number_str)//2)):
        return True
    
    
