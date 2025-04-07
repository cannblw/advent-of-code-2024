left, right = [], []

with open("input", "r") as f:
    for line in f.readlines():
        tokens = line.split()

        left.append(tokens[0])
        right.append(tokens[1])

left.sort()
right.sort()

acc = 0

for i, l in enumerate(left):
    acc = acc + abs(int(right[i]) - int(l))

print(acc)
