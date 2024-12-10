import sys


def parse_grid(intext: str) -> list[list[int]]:
    return [[int(x) for x in line] for line in intext.split("\n")]


def get_neighbors(
    tile: tuple[int, int], height: int, width: int
) -> list[tuple[int, int]]:
    adjs = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    neighbors = []
    for adj in adjs:
        nr = tile[0] + adj[0]
        nc = tile[1] + adj[1]
        if 0 <= nr < height and 0 <= nc < width:
            neighbors.append((nr, nc))
    return neighbors


def part1(intext: str) -> int:
    grid = parse_grid(intext)

    height = len(grid)
    width = len(grid[0])

    trailheads: list[tuple[int, int]] = []

    for r in range(height):
        for c in range(width):
            if grid[r][c] == 0:
                trailheads.append((r, c))

    def dfs(cur: int, tile: tuple[int, int], visited: set[tuple[int, int]]) -> int:
        if cur == 9:
            return 1
        neighbors = get_neighbors(tile, height, width)
        total = 0
        for n in neighbors:
            if n in visited:
                continue
            if grid[n[0]][n[1]] == cur + 1:
                visited.add(n)
                total += dfs(cur + 1, n, visited)

        return total

    total = 0
    for th in trailheads:
        score = dfs(0, th, set())
        total += score

    return total


def part2(intext: str) -> int:
    grid = parse_grid(intext)

    height = len(grid)
    width = len(grid[0])

    trailheads: list[tuple[int, int]] = []

    for r in range(height):
        for c in range(width):
            if grid[r][c] == 0:
                trailheads.append((r, c))

    def dfs(cur: int, tile: tuple[int, int], visited: set[tuple[int, int]]) -> int:
        if cur == 9:
            return 1
        neighbors = get_neighbors(tile, height, width)
        total = 0
        for n in neighbors:
            if n in visited:
                continue
            if grid[n[0]][n[1]] == cur + 1:
                visited.add(n)
                total += dfs(cur + 1, n, visited)
                visited.remove(n)

        return total

    total = 0
    for th in trailheads:
        score = dfs(0, th, set())
        total += score

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
