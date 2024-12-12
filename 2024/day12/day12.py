import sys
from collections import defaultdict


def parse_grid(intext: str) -> list[list[str]]:
    return [[char for char in line] for line in intext.split("\n")]


def inbounds(coord: tuple[int, int], height: int, width: int) -> bool:
    return 0 <= coord[0] < height and 0 <= coord[1] < width


def get_neighbors(
    coord: tuple[int, int], height: int, width: int
) -> set[tuple[int, int]]:
    neighbors = set()
    cr, cc = coord
    for dr, dc in ((-1, 0), (1, 0), (0, -1), (0, 1)):
        n = (cr + dr, cc + dc)
        if inbounds(n, height, width):
            neighbors.add(n)
    return neighbors


def part1(intext: str) -> int:
    grid = parse_grid(intext)

    height = len(grid)
    width = len(grid[0])

    visited: set[tuple[int, int]] = set()

    def dfs(coord: tuple[int, int], char: str, region: set[tuple[int, int]]):
        visited.add(coord)
        region.add(coord)
        for nr, nc in get_neighbors(coord, height, width):
            if (
                (nr, nc) not in visited
                and (nr, nc) not in region
                and grid[nr][nc] == char
            ):
                dfs((nr, nc), char, region)

    total = 0
    for r in range(height):
        for c in range(width):
            if (r, c) in visited:
                continue
            region = set()
            dfs((r, c), grid[r][c], region)

            area = len(region)

            # edges for each tile = 4 - num of adj tiles in same region
            # perimeter = num edges
            per = sum(
                4 - len(region.intersection(get_neighbors(x, height, width)))
                for x in region
            )

            # print(grid[r][c], region, area, per)
            total += area * per

    return total


def part2(intext: str) -> int:
    grid = parse_grid(intext)

    height = len(grid)
    width = len(grid[0])

    visited: set[tuple[int, int]] = set()

    def dfs(coord: tuple[int, int], char: str, region: set[tuple[int, int]]):
        visited.add(coord)
        region.add(coord)
        for nr, nc in get_neighbors(coord, height, width):
            if (
                (nr, nc) not in visited
                and (nr, nc) not in region
                and grid[nr][nc] == char
            ):
                dfs((nr, nc), char, region)

    total = 0
    for r in range(height):
        for c in range(width):
            if (r, c) in visited:
                continue
            region = set()
            dfs((r, c), grid[r][c], region)

            area = len(region)

            # num sides = num corners
            # real corners exist where 1 or 3 tiles touch corner
            # edge case for mobius where 2 tiles touch corner
            reg_cors = defaultdict(set[tuple[int, int]])
            for xr, xc in region:
                for dr, dc in ((0, 0), (0, 1), (1, 0), (1, 1)):
                    reg_cors[(xr + dr, xc + dc)].add((dr, dc))
            sides = 0
            for origins in reg_cors.values():
                if len(origins) == 2:
                    a, b = origins
                    if (a[0] + b[0], a[1] + b[1]) == (1, 1):  # mobius corner case
                        sides += 2
                elif len(origins) % 2 == 1:
                    sides += 1

            # print(grid[r][c], region, area, sides)
            total += area * sides

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

# this code is pretty gross
