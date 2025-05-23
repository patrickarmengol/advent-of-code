def prio(c: str) -> int:
    return ord(c) - \
        96 if c.islower() else ord(c) - 64 + 26


def part1() -> None:
    total = 0
    with open('day03/input.txt') as f:
        sacks: list[str] = f.read().rstrip('\n').split('\n')
        for sack in sacks:
            midpoint = len(sack)//2
            comp1, comp2 = sack[:midpoint], sack[midpoint:]
            intersect = set(comp1) & set(comp2)
            interchar = intersect.pop()
            total += prio(interchar)
    print(total)


def part2() -> None:
    total = 0
    with open('day03/input.txt') as f:
        sacks: list[str] = f.read().rstrip('\n').split('\n')
        groups = [sacks[i:i+3] for i in range(0, len(sacks), 3)]
        for group in groups:
            intersect = set(group[0]) & set(group[1]) & set(group[2])
            interchar = intersect.pop()
            total += prio(interchar)
    print(total)


if __name__ == '__main__':
    part1()
    part2()
