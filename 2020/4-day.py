import sys, timeit

# Remove \n from string
def chomp(s):
    return s[:-1] if s.endswith('\n') else s

# Parse file and return a list
# Algo :
# Create 3 buffers :
#   - listPassClean, contains passport data non normalized
#   - listBuff, temporary list work
#   - listFinal, contains passport data normalized as one field by slot list
# Exemple In : 
# eyr:2035
# byr:1988 hgt:193cm
# iyr:2028 cid:128 hcl:#18171d ecl:utc pid:9743739773
# Out : ['eyr:2035', 'byr:1988', 'hgt:193cm', 'iyr:2028', 'cid:128', 'hcl:#18171d', 'ecl:utc', 'pid:9743739773']
def parseFile(listPassport):
    listPassClean = []
    listBuff = []
    
    # add passport information on multiple line into dedicated list
    # then add this new list into listPassClean (data non normalized)
    # Exemple :
    # IN :
    # eyr:2035
    # byr:1988 hgt:193cm
    # iyr:2028 cid:128 hcl:#18171d ecl:utc pid:9743739773
    # OUT : ['eyr:2035', 'byr:1988 hgt:193cm', 'iyr:2028 cid:128 hcl...9743739773']
    for line in listPassport:
        if line == '':
            listPassClean.append(listBuff)
            listBuff = []
        else:
            listBuff.append(line)
    
    # NOTE : Because the file is not ending py "\n"
    # we must add the content of the buffer at the end of the loop
    # Otherwise last line of file will not be handle.
    listPassClean.append(listBuff)

    listBuff = []
    listFinal = []
    
    # normalized each slot of the "non-normalized" list
    # Exemple : 
    # IN  : ['eyr:2035', 'byr:1988 hgt:193cm', 'iyr:2028 cid:128 hcl...9743739773']
    # OUT : ['eyr:2035', 'byr:1988', 'hgt:193cm', 'iyr:2028', 'cid:128', 'hcl:#18171d', 'ecl:utc', 'pid:9743739773']
    for _list in listPassClean:
        listBuff = []
        for string in _list:
            if " " in string:
                stringSplit = string.split(" ")
                for elm in stringSplit:
                    listBuff.append(elm)
            else:
                listBuff.append(string)
        listFinal.append(listBuff)
            

    return listFinal

# return a list which contains only validated passport data
def checkField(list2DPassport, listFieldRequired, listFieldOptional):
    listClean = []
    for _list in list2DPassport:
        isOk = True
        # quick check to eliminate list
        if len(_list) < len(listFieldRequired):
            continue
        else:
            listBuff = []
            # parse and insert all field of passport into a list
            for field_list in _list:
                f = field_list.split(":")[0]
                listBuff.append(f)

            # check if 
            for fieldRequired in listFieldRequired:
                if fieldRequired not in listBuff:
                    isOk = False
                    break
                else:
                    continue
            
            if isOk :
                listClean.append(_list)

    return listClean

f=sys.argv[1]
listFieldRequired = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
listFieldOptional = ["cid"]

lClean = []
with open(f,'r') as it:
    lines = it.readlines()
    for l in lines:
        lClean.append(chomp(l))
    del lines

lClean = parseFile(lClean)
lClean = checkField(lClean, listFieldRequired, listFieldOptional)

print(len(lClean))