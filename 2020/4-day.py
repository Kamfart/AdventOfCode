import sys, timeit

tic=timeit.default_timer()

# INSTRUCTION PART 1 : The automatic passport scanners are slow because they're having trouble detecting which passports have all required fields.
# Count the number of valid passports - those that have all required fields. Treat cid as optional. In your batch file, how many passports are valid?
# Required field :  byr (Birth Year) ; iyr (Issue Year) ; eyr (Expiration Year) ; hgt (Height) ; hcl (Hair Color) ; ecl (Eye Color) ; pid (Passport ID) ; cid (Country ID)

# INSTRUCTION PART 2 : You can continue to ignore the cid field, but each other field has strict rules about what values are valid for automatic validation.
# Your job is to count the passports where all required fields are both present and valid
# Rules :  
    # byr (Birth Year) - four digits; at least 1920 and at most 2002.
    # iyr (Issue Year) - four digits; at least 2010 and at most 2020.
    # eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
    # hgt (Height) - a number followed by either cm or in:
    #   If cm, the number must be at least 150 and at most 193.
    #   If in, the number must be at least 59 and at most 76.
    # hcl (Hair Color) - a '#' followed by exactly six characters 0-9 or a-f.
    # ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
    # pid (Passport ID) - a nine-digit number, including leading zeroes.

# Remove \n from string
def chomp(s):
    return s[:-1] if s.endswith('\n') else s

# Parse file and return a list
# Exemple :
# IN : 
# eyr:2035
# byr:1988 hgt:193cm
# iyr:2028 cid:128 hcl:#18171d ecl:utc pid:9743739773
# OUT : ['eyr:2035', 'byr:1988', 'hgt:193cm', 'iyr:2028', 'cid:128', 'hcl:#18171d', 'ecl:utc', 'pid:9743739773']
def parseFile(listPassport):
    listPassClean = []
    listBuff = []
    
    # add passport information on multiple line into a buffer list
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
    
    # NOTE : Because the file is not ending by "\n"
    # we must add the content of the buffer at the end of the loop
    # Otherwise the last line of the file will not be handle.
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

# check if a field has the required length and if it is between a given interval
def controlYear(field, lenDigit, start, end):
    if len(field) != lenDigit or int(field) < start or int(field) > end:
        return False
    else:
        return True

# check if a field has the required unit and if it is between a given interval
# param : min/max for cm and inches
def controlHeight(field, minCm, maxCm, minIn, maxIn):
    value=field[:-2]
    unit=field[-2:]
    if unit != "in" and unit != "cm":
        return False
    elif unit == "in" and (int(value) < minIn or int(value) > maxIn):
        return False
    elif unit == "cm" and (int(value) < minCm or int(value) > maxCm):
        return False
    else:
        return True

# check if a field has the required character and if it composed by hexadecimal value only
def controlHairColor(field):
    valueHexa=["0","1","2","3","4","5","6","7","8","9","a","b","c","d","e","f"]
    if field[0] != "#":
        return False
    else:
        for c in field[1:]:
            if c not in valueHexa:
                return False
        return True

# check if a field is include in a defined list
def controlEyeColor(field):
    color = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]
    if field not in color:
        return False
    else:
        return True

# check if a field has the required length and if it's a digit
def controlPassportID(field, lenID):
    if len(field) != lenID or not field.isdigit():
        return False
    else:
        return True

# Main controle : 
# Algo :
# Iterate over 2DList, each _list of the 1st iteration is iterate in another loop
# Each field of the 2nd itaration is parse in key and value variable
# Each key has if own test, if one test is false, break the current loop
# otherwise add the _list of the 1st iteration into a buffer list
# Return the controled list
def controlField(list2DPassport):
    listClean = []

    for _list in list2DPassport:
        isOk = True

        for kv in _list:
            key = kv.split(":")[0]
            value = kv.split(":")[1]
            if key == "byr":
                if not controlYear(value, 4, 1920, 2002):
                    isOk = False
                    break
            elif key == "iyr":
                if not controlYear(value, 4, 2010, 2020):
                    isOk = False
                    break
            elif key == "eyr":
                if not controlYear(value, 4, 2020, 2030):
                    isOk = False
                    break
            elif key == "hgt":
                if not controlHeight(value, 150, 193, 59, 76):
                    isOk = False
                    break
            elif key == "hcl":
                if not controlHairColor(value):
                    isOk = False
                    break
            elif key == "ecl":
                if not controlEyeColor(value):
                    isOk = False
                    break
            elif key == "pid":
                if not controlPassportID(value, 9):
                    isOk = False
                    break
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
lClean = controlField(lClean)

print(len(lClean))

toc=timeit.default_timer()
print(toc - tic)