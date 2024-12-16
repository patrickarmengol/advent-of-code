import sys
import re
from collections import defaultdict


def parse_robots(intext: str) -> list[tuple[int, int, int, int]]:
    robots = []
    for px, py, vx, vy in re.findall(r"p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)", intext):
        robots.append((int(px), int(py), int(vx), int(vy)))
    return robots


def part1(intext: str) -> int:
    robots = parse_robots(intext)
    steps = 100
    height = 103
    width = 101

    dests: defaultdict[tuple[int, int], int] = defaultdict(int)  # could be counter
    for px, py, vx, vy in robots:
        # dest = start + n_steps * velocity
        dest = ((px + (steps * vx)) % width, (py + (steps * vy)) % height)
        dests[dest] += 1

    quad = defaultdict(int)
    for (dx, dy), count in dests.items():
        if dx < width // 2:
            if dy < height // 2:
                quad["nw"] += count
            elif dy > height // 2:
                quad["sw"] += count
        elif dx > width / 2:
            if dy < height // 2:
                quad["ne"] += count
            elif dy > height // 2:
                quad["se"] += count
    res = 1
    for q in quad.values():
        res *= q
    return res


def print_robots(
    robots: defaultdict[tuple[int, int], int], height: int, width: int
) -> None:
    for r in range(height):
        for c in range(width):
            if (c, r) in robots:
                print(robots[(c, r)], end="")
            else:
                print(".", end="")
        print()


def is_dense_somewhere(robots: defaultdict[tuple[int, int], int]) -> bool:
    sector_counts = defaultdict(int)
    n_sectors = 8

    for dx, dy in robots:
        sector_counts[(dx // n_sectors, dy // n_sectors)] += 1

    return any(count > 20 for count in sector_counts.values())


def part2(intext: str) -> int:
    robots = parse_robots(intext)
    height = 103
    width = 101

    i = 0
    while True:
        # print(i)
        dests: defaultdict[tuple[int, int], int] = defaultdict(int)  # could be counter
        for px, py, vx, vy in robots:
            # dest = start + n_steps * velocity
            dest = ((px + (i * vx)) % width, (py + (i * vy)) % height)
            dests[dest] += 1

        if is_dense_somewhere(dests):
            print_robots(dests, height, width)
            return i
        i += 1


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
