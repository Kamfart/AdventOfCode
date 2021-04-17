import sys, re

class Luggage:
    def __init__(self, id, color, content):
        self.id = id
        self.color = color
        self.content = content

# c = [(4,1),(2,2)]
# l1 = Luggage(0, "blue", [(4,1),(2,2)])

# print(l1.content[0][0])

def parseFile(listRules, listLuggages):
    for line in listRules:
        id = getLastIdLuggage(listLuggages)
        color = parseColorLuggage(line)
        content = parseContentLuggage(line)
        luggage = Luggage(id, color, content)
        listLuggages.append(luggage)

def getLastIdLuggage(listLuggages):
    size = len(listLuggages)
    if size == 0:
        return 0
    else:
        return size + 1

def parseColorLuggage(line):
    return line.split("contain")[0].strip()

def parseContentLuggage(line):
    content = []
    buff = line.split("contain")[1].strip()
    if "," in buff:
        buff = buff.split(",")
        for c in buff:
            buff_2 = c.strip().split(" ", 1)
            buff_2[1] = buff_2[1].strip(".")
            content.append(buff_2)
    else:
        #content.append(buff.strip().split(" ", 1))
        buff = buff.strip().split(" ", 1)
        buff[1] = buff[1].strip(".")
        content.append(buff)
    
    return content

def cleanFiles(lines, pattern, repl):
    newLine=[]
    for line in lines:
        newLine.append(re.sub(pattern, repl, line))

    return newLine

def getContentLuggage(luggage):
    return luggage.content

def getColorLuggage(luggage):
    return luggage.color

def solve1(listLuggages, luggage):
    possibleLuggages = []
    for lug in listLuggages:
        content = getContentLuggage(lug)
        for bag in content:
            bagName = bag[1]
            if bagName == luggage:
                possibleLuggages.append(lug)
            else:
                continue
    
    for possibleLug in possibleLuggages:
        color = getColorLuggage(possibleLug)
        for lug in listLuggages:
            if lug in possibleLuggages:
                continue
            else:
                content = getContentLuggage(lug)
                for bag in content:
                    bagName = bag[1]
                    if bagName == color:
                        possibleLuggages.append(lug)
                    else:
                        continue
        
    return possibleLuggages

f=sys.argv[1]
listLuggages=[]
with open(f,'r') as _file:
    lines = _file.readlines()
    lines = cleanFiles(lines, 'bags', 'bag')
    parseFile(lines, listLuggages)  
    del lines

luggage = "shiny gold bag"

listValidLuggage = solve1(listLuggages, luggage)

print('Solution part 1 : {:d}'.format(len(listValidLuggage)))