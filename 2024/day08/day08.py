from itertools import combinations
import sys
from collections import defaultdict


def parse_grid(intext: str) -> list[list[str]]:
    return [[char for char in line] for line in intext.split("\n")]


def find_antennas(grid: list[list[str]]) -> defaultdict[str, list[tuple[int, int]]]:
    d = defaultdict(list[tuple[int, int]])
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] != ".":
                d[grid[r][c]].append((r, c))
    return d


def print_grid(grid: list[list[str]], antis: set[tuple[int, int]]):
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if (r, c) in antis and grid[r][c] == ".":
                print("#", end="")
            else:
                print(grid[r][c], end="")
        print()


def inbounds(p: tuple[int, int], height: int, width: int) -> bool:
    return 0 <= p[0] < height and 0 <= p[1] < width


def part1(intext: str) -> int:
    grid = parse_grid(intext)

    height = len(grid)
    width = len(grid[0])

    antennas = find_antennas(grid)

    antinodes: set[tuple[int, int]] = set()

    for freq in antennas:
        combos = combinations(antennas[freq], 2)
        for a, b in combos:
            diff_r = a[0] - b[0]
            diff_c = a[1] - b[1]

            anti_a = (a[0] + diff_r, a[1] + diff_c)
            anti_b = (b[0] - diff_r, b[1] - diff_c)
            if inbounds(anti_a, height, width):
                antinodes.add(anti_a)
            if inbounds(anti_b, height, width):
                antinodes.add(anti_b)

    return len(antinodes)


def part2(intext: str) -> int:
    grid = parse_grid(intext)

    height = len(grid)
    width = len(grid[0])

    antennas = find_antennas(grid)

    antinodes: set[tuple[int, int]] = set()

    for freq in antennas:
        combos = combinations(antennas[freq], 2)
        for a, b in combos:
            diff_r = a[0] - b[0]
            diff_c = a[1] - b[1]

            # a-side
            i = 0
            while inbounds(
                anti := (a[0] + (diff_r * i), a[1] + (diff_c * i)), height, width
            ):
                antinodes.add(anti)
                i += 1
            # b-side
            i = 0
            while inbounds(
                anti := (b[0] - (diff_r * i), b[1] - (diff_c * i)), height, width
            ):
                antinodes.add(anti)
                i += 1

        # print_grid(grid, antinodes)
        # print()
    return len(antinodes)


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
