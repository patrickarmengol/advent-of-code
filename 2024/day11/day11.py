import sys
from collections import defaultdict


def parse_stones(intext: str) -> list[str]:
    return intext.split()


def part1(intext: str) -> int:
    stones = parse_stones(intext)

    for i in range(25):
        new_stones = []
        for stone in stones:
            if stone == "0":
                new_stones.append("1")
            elif len(stone) % 2 == 0:
                a, b = stone[: len(stone) // 2], str(int(stone[len(stone) // 2 :]))
                new_stones.append(a)
                new_stones.append(b)
            else:
                new_stones.append(str(int(stone) * 2024))
        stones = new_stones

    return len(stones)


def blink(s: str) -> list[str]:
    bs = []
    if s == "0":
        bs.append("1")
    elif len(s) % 2 == 0:
        bs.append(s[: len(s) // 2])
        bs.append(str(int(s[len(s) // 2 :])))
    else:
        bs.append(str(int(s) * 2024))
    return bs


def part2(intext: str) -> int:
    stones = parse_stones(intext)

    # counter for stones
    c = defaultdict(int)
    for s in stones:
        c[s] += 1

    for _ in range(75):
        nc = defaultdict(int)
        # iterate through cur counter (list of stones, counts)
        for stone, count in c.items():
            # for each stone, blink; new count of stone is existing of stone * newly blinked of stone
            for s in blink(stone):
                nc[s] += count
        c = nc

    # sol would be sum of count of all stones
    return sum(c.values())


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

# part1 solution is naive
# for part2 i initially didn't recognize a pattern, but thought about how each unique stone increments
# some others' solutions involve recursive cached functions
