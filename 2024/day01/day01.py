import sys
from collections import Counter


def parse_lists(intext: str) -> tuple[list[int], list[int]]:
    a, b = [], []
    for line in intext.split("\n"):
        c, d = line.split()
        a.append(int(c))
        b.append(int(d))
    return (a, b)


def part1(intext: str) -> int:
    a, b = parse_lists(intext)

    a.sort()
    b.sort()

    count = 0
    for x, y in zip(a, b):
        diff = abs(x - y)
        count += diff

    return count


def part2(intext: str) -> int:
    a, b = parse_lists(intext)

    count = 0
    bc = Counter(b)
    for x in a:
        if x in bc:
            count += x * bc[x]

    return count


# read input and call part functions
def main():
    if len(sys.argv) < 3:
        print("usage: python day<NN>.py <part> <infile>")
        sys.exit(1)
    part = sys.argv[1]
    infile = sys.argv[2]
    intext = open(infile).read().strip()

    if part in ("0", "1"):
        print(part1(intext))
    if part in ("0", "2"):
        print(part2(intext))


if __name__ == "__main__":
    main()
