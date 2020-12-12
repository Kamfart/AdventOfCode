# Instruction part 1 : Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.
# Exemple :     199 + 1821 = 2020 ; 199 * 1821 = key 

# Instruction part 2 : Same, with 3 numbers.
import sys, timeit

tic=timeit.default_timer()

f=sys.argv[1]

# Remove \n from string
def chomp(s):
    return s[:-1] if s.endswith('\n') else s

# Algo :
# 1) convert all elem from list into int
# 2) loop over the list twice in the same time to compare 2 value
# 3) create buff_index value to skip if the same value is tested
def solve1(listValue, valueWanted):
    listValue = [int(elm) for elm in listValue]
    indexI = 0
    for i in listValue:
        indexJ = 0
        for j in listValue:
            if indexI == indexJ:
                continue
            elif (i + j) == valueWanted:
                return i * j
                
            else:
                indexJ+=1
        indexI+=1
# Algo :
# Same than solve1 with 3 loops
def solve2(listValue, valueWanted):
    listValue = [int(elm) for elm in listValue]
    indexI = 0
    for i in listValue:
        indexJ = 0
        for j in listValue:
            IndexK = 0
            for k in listValue:
                if indexI == indexJ == IndexK:
                    continue
                elif (i + j + k) == valueWanted:
                    return i * j * k
                else:
                    indexJ+=1
        indexI+=1



lClean = []
# Add file into the list and clean /n of each slot
with open(f,'r') as it:
    lines = it.readlines()
    for l in lines:
        lClean.append(chomp(l))
    del lines

print(solve1(lClean, 2020))
print(solve2(lClean, 2020))

toc=timeit.default_timer()
print(toc - tic)