import sys
import networkx as nx


def parse_grid(intext: str) -> list[list[str]]:
    return [[char for char in line] for line in intext.split("\n")]


def find_source_end(grid: list[list[str]]) -> tuple[tuple[int, int], tuple[int, int]]:
    s = (0, 0)
    e = (0, 0)
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == "S":
                s = (r, c)
            elif grid[r][c] == "E":
                e = (r, c)
    return (s, e)


def parse_graph(
    grid: list[list[str]], s: tuple[int, int, str], e: tuple[int, int]
) -> nx.Graph:
    G = nx.DiGraph()
    height = len(grid)
    width = len(grid[0])
    for r in range(height):
        for c in range(width):
            # walls are not nodes
            if grid[r][c] == "#":
                continue
            i = 0
            sides = "NESWN"
            adjs = {
                "N": (-1, 0),
                "E": (0, 1),
                "S": (1, 0),
                "W": (0, -1),
            }
            # for each direction
            for i in range(len(sides) - 1):
                # add edge w=1000 between directions on same tile
                if (r, c) != e:
                    G.add_edge((r, c, sides[i]), (r, c, sides[i + 1]), weight=1000)
                    G.add_edge((r, c, sides[i + 1]), (r, c, sides[i]), weight=1000)
                # add edge w=1 between adjacent edges with same dir
                ar, ac = adjs[sides[i]]
                nr, nc = (r + ar, c + ac)
                if 0 <= nr < height and 0 <= nc < width and grid[nr][nc] != "#":
                    G.add_edge(
                        (r, c, sides[i]),
                        (nr, nc, sides[i]) if (nr, nc) != e else (nr, nc),
                        weight=1,
                    )
                i += 1
    return G


def print_grid(grid: list[list[str]]):
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            print(grid[r][c], end="")
        print()


def part1(intext: str) -> int:
    grid = parse_grid(intext)
    s, e = find_source_end(grid)
    s = (s[0], s[1], "E")
    G = parse_graph(grid, s, e)

    # get path and distance with dijkstra's
    p, d = nx.single_source_dijkstra_path(G, s, e)

    # fill in grid with path for printing
    for step in p:
        if len(step) == 3:
            grid[step[0]][step[1]] = step[2]
    print_grid(grid)

    return d


def part2(intext: str) -> int:
    grid = parse_grid(intext)
    s, e = find_source_end(grid)
    s = (s[0], s[1], "E")
    G = parse_graph(grid, s, e)

    # get all shortest paths with dijkstra's
    ps = nx.all_shortest_paths(G, s, e, weight="weight")

    # fill in grid with unique tiles from all paths for printing
    bests = set()
    for p in ps:
        for tile in p:
            bests.add((tile[0], tile[1]))
    for br, bc in bests:
        grid[br][bc] = "O"
    print_grid(grid)

    return len(bests)


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
