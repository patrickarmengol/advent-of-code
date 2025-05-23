def lines(filepath: str) -> list[str]:
    with open(filepath) as f:
        return f.read().rstrip('\n').split('\n')


def part1() -> None:
    total = 0
    pairs = lines(problem_input)
    for pair in pairs:
        elf1, elf2 = pair.split(',')
        elf1_min, elf1_max = [int(n) for n in elf1.split('-')]
        elf2_min, elf2_max = [int(n) for n in elf2.split('-')]
        if (elf1_min <= elf2_min and elf1_max >= elf2_max) \
                or (elf2_min <= elf1_min and elf2_max >= elf1_max):
            total += 1
    print(total)


def part2() -> None:
    total = 0
    pairs = lines(problem_input)
    for pair in pairs:
        elf1, elf2 = pair.split(',')
        elf1_min, elf1_max = [int(n) for n in elf1.split('-')]
        elf2_min, elf2_max = [int(n) for n in elf2.split('-')]
        elf1_set = set(range(int(elf1_min), int(elf1_max) + 1))
        elf2_set = set(range(int(elf2_min), int(elf2_max) + 1))
        if (elf1_set & elf2_set):
            total += 1
    print(total)


if __name__ == '__main__':
    problem_input = 'day04/input.txt'
    part1()
    part2()
