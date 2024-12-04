from itertools import combinations
import sys


def part1(intext: str) -> int:
    grid = [[c for c in line] for line in intext.split("\n")]

    height = len(grid)
    width = len(grid[0])

    word = "XMAS"

    def recurse(r: int, c: int, i: int, dir: tuple[int, int]) -> int:
        if i == len(word):
            return 1
        if not (0 <= r < height and 0 <= c < width and grid[r][c] == word[i]):
            return 0
        return recurse(r + dir[0], c + dir[1], i + 1, dir)

    total = 0
    for r in range(height):
        for c in range(width):
            for dir in (
                (0, 1),
                (1, 0),
                (-1, 0),
                (0, -1),
                (-1, -1),
                (-1, 1),
                (1, -1),
                (1, 1),
            ):
                total += recurse(r, c, 0, dir)
    return total


def part2(intext: str) -> int:
    grid = [[c for c in line] for line in intext.split("\n")]

    height = len(grid)
    width = len(grid[0])

    starts = [(r, c) for r in range(height) for c in range(width) if grid[r][c] == "A"]

    total = 0
    for r, c in starts:
        adj_str = "".join(
            [
                grid[r + dr][c + dc]
                for dr, dc in [(-1, -1), (-1, 1), (1, 1), (1, -1)]
                if (0 <= r + dr < height and 0 <= c + dc < width)
            ]
        )
        if adj_str in ("MMSS", "MSSM", "SSMM", "SMMS"):
            total += 1

    return total


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
