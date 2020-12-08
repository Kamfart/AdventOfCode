# Instruction : Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.
# Exemple :     199 + 1821 = 2020 ; 199 * 1821 = key 
import sys, timeit

tic=timeit.default_timer()

f=sys.argv[1]

def chomp(s):
    return s[:-1] if s.endswith('\n') else s

def solve1(listValue, valueWanted):
    indexI = 0
    for i in listValue:
        indexJ = 0
        for j in listValue:
            if indexI == indexJ:
                indexJ+=1
            elif (int(i) + int(j)) == valueWanted:
                return (int(i) * int(j))
                
            else:
                indexJ+=1  
        indexI+=1

def solve2(listValue, valueWanted):
    listValue = [int(elm) for elm in listValue]
    indexI = 0
    for i in listValue:
        indexJ = 0
        for j in listValue:
            if indexI == indexJ:
                indexJ+=1
            elif (i + j) == valueWanted:
                return i * j
                
            else:
                indexJ+=1
        indexI+=1

lClean = []
with open(f,'r') as it:
    lines = it.readlines()
    for l in lines:
        lClean.append(chomp(l))
    del lines

print(solve2(lClean, 2020))

toc=timeit.default_timer()
print(toc - tic)