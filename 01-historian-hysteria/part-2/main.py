left = []
numbersAndOccurrences = {}

with open("input", "r") as f:
    for line in f.readlines():
        tokens = line.split()

        t1, t2 = int(tokens[0]), int(tokens[1])

        left.append(t1)

        if t2 not in numbersAndOccurrences:
            numbersAndOccurrences[t2] = 0

        numbersAndOccurrences[t2] = numbersAndOccurrences[t2] + 1


similarity = 0

for item in left:
    if item in numbersAndOccurrences:
        similarity = similarity + (item * numbersAndOccurrences[item])

print(similarity)
