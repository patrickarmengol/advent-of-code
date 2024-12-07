import sys

dir_map = {
    "<": (0, -1),
    ">": (0, 1),
    "^": (-1, 0),
    "V": (1, 0),
}
turn_map = {
    "<": "^",
    ">": "V",
    "^": ">",
    "V": "<",
}

reverse_turn_map = {
    "<": "V",
    ">": "^",
    "^": "<",
    "V": ">",
}


def simple_grid(grid: list[list[str]]):
    for r in range(len(grid)):
        for c in range(len(grid[r])):
            print(grid[r][c], end="")
        print()


def print_grid(grid: list[list[str]], vv: set[tuple[tuple[int, int], str]]):
    d = {vp: vd for vp, vd in vv}
    for r in range(len(grid)):
        for c in range(len(grid[r])):
            if (r, c) in d:
                print(d[(r, c)], end="")
            else:
                print(grid[r][c], end="")
        print()


def parse_grid(intext: str) -> list[list[str]]:
    return [[char for char in line] for line in intext.split("\n")]


def find_start(grid: list[list[str]]) -> tuple[tuple[int, int], str]:
    for r in range(len(grid)):
        for c in range(len(grid[r])):
            if grid[r][c] in "<>^V":
                return ((r, c), grid[r][c])
    raise ValueError("guard not found in starting grid")


def part1(intext: str) -> int:
    grid = parse_grid(intext)

    height = len(grid)
    width = len(grid[0])

    guard = find_start(grid)

    visited: set[tuple[int, int]] = set()

    while True:
        cur_pos = guard[0]
        cur_dir = guard[1]
        visited.add(cur_pos)
        grid[cur_pos[0]][cur_pos[1]] = cur_dir

        # sleep(0.01)
        # simple_grid(grid)
        # print()

        next_pos = (cur_pos[0] + dir_map[cur_dir][0], cur_pos[1] + dir_map[cur_dir][1])
        next_dir = cur_dir
        if not (0 <= next_pos[0] < height and 0 <= next_pos[1] < width):
            break
        if grid[next_pos[0]][next_pos[1]] == "#":
            next_pos = cur_pos
            next_dir = turn_map[cur_dir]

        guard = (next_pos, next_dir)

    simple_grid(grid)
    return len(visited)


def part2(intext: str) -> int:
    grid = parse_grid(intext)

    height = len(grid)
    width = len(grid[0])

    initial_guard = find_start(grid)

    def simulate(
        g: tuple[tuple[int, int], str], v: set[tuple[tuple[int, int], str]]
    ) -> tuple[set[tuple[tuple[int, int], str]], bool]:
        while True:
            cur_pos = g[0]
            cur_dir = g[1]

            if (cur_pos, cur_dir) in v:
                return v, False
            v.add((cur_pos, cur_dir))

            next_pos = (
                cur_pos[0] + dir_map[cur_dir][0],
                cur_pos[1] + dir_map[cur_dir][1],
            )
            next_dir = cur_dir
            if not (0 <= next_pos[0] < height and 0 <= next_pos[1] < width):
                return v, True
            if grid[next_pos[0]][next_pos[1]] in "#O":
                next_pos = cur_pos
                next_dir = turn_map[cur_dir]

            g = (next_pos, next_dir)

    visited, _ = simulate(initial_guard, set())

    rocks = set()
    for vis in visited:
        if vis[0] in rocks:
            continue
        vis_r, vis_c = vis[0]
        grid[vis_r][vis_c] = "O"
        _, ok = simulate(initial_guard, set())
        if not ok:
            rocks.add((vis_r, vis_c))
        grid[vis_r][vis_c] = "."

    return len(rocks)


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

# this code is not great
# and this solution involves brute forcing (which i thought i needed to avoid)
#
# i'm pretty sure there's a neat solution that solves it in one pass:
#
# simulate as usual
# at each turning point, go back until obstacle behind, marking tiles as traps
# if found an obstacle on left of backtrack, also backtrack on that track
# continuing the main simulation,
# if the guard hits a trap (where dir of trap is clockwise of guard),
# then the guards next position (in front) would be a new viable block point
