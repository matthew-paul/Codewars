import itertools
def permutations(s):
    return set(''.join(x) for x in itertools.permutations(s, r=len(s)))