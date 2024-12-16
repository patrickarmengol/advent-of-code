import sys


def parse_grid_and_moves(intext: str) -> tuple[list[list[str]], str]:
    map, instr = intext.split("\n\n")
    grid = [[char for char in line] for line in map.split("\n")]
    moves = "".join(instr.split("\n"))
    return grid, moves


def find_robot(grid: list[list[str]]) -> tuple[int, int]:
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == "@":
                return (r, c)
    assert False


def resize_grid(grid: list[list[str]]) -> list[list[str]]:
    new_grid = [["A" for _ in range(len(grid[0]) * 2)] for _ in range(len(grid))]
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            new_grid[r][2 * c] = grid[r][c] if grid[r][c] in "#.@" else "["
            new_grid[r][2 * c + 1] = (
                grid[r][c] if grid[r][c] in "#." else "." if grid[r][c] == "@" else "]"
            )
    return new_grid


def find_boxes(grid: list[list[str]]) -> set[tuple[int, int]]:
    s = set()
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] in "[O":
                s.add((r, c))
    return s


move_map = {
    "<": (0, -1),
    ">": (0, 1),
    "^": (-1, 0),
    "v": (1, 0),
}


def print_grid(grid: list[list[str]]):
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            print(grid[r][c], end="")
        print()


def part1(intext: str) -> int:
    grid, moves = parse_grid_and_moves(intext)
    rr, rc = find_robot(grid)

    for move in moves:
        # print_grid(grid)
        # print(move)
        dr, dc = move_map[move]
        nr, nc = (rr + dr, rc + dc)
        if grid[nr][nc] == "#":
            continue
        elif grid[nr][nc] == ".":
            grid[rr][rc] = "."
            grid[nr][nc] = "@"
            rr, rc = nr, nc
        else:
            br, bc = nr, nc
            while grid[br][bc] == "O":
                br, bc = (br + dr, bc + dc)
            if grid[br][bc] == ".":  # empty space for balls
                grid[br][bc] = "O"
                grid[rr][rc] = "."
                grid[nr][nc] = "@"
                rr, rc = nr, nc
            else:  # nowhere to move balls
                continue

    total = 0
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == "O":
                total += (100 * r) + c
    return total


def move_box(
    box: tuple[int, int],
    dir: tuple[int, int],
    grid: list[list[str]],
    boxes: set[tuple[int, int]],
) -> set[tuple[tuple[int, int], tuple[int, int]]]:
    br, bc = box
    dr, dc = dir
    if (dr, dc) in ((0, -1), (0, 1)):  # horizontal
        ar, ac = (br, bc + (2 * dc))  # adj box
        tr, tc = (br, bc - 1 if dc == -1 else bc + 2)  # adj tile
        if (ar, ac) in boxes:  # depends on other boxes
            rm = move_box((ar, ac), dir, grid, boxes)
            if rm:
                change = ((br, bc), (br, bc + dc))
                return set([change]).union(rm)
            else:
                return set()
        elif grid[tr][tc] == "#":  # can't move
            return set()
        else:  # can move
            change = ((br, bc), (br, bc + dc))
            return set([change])
    else:  # vertical
        adjs = set((br + dr, bc + gc) for gc in (0, 1))
        ass = {}
        for ar, ac in adjs:
            if grid[ar][ac] == "]":
                ac = ac - 1
            if (ar, ac) in boxes:
                rm = move_box((ar, ac), dir, grid, boxes)
                if rm:
                    change = ((br, bc), (br + dr, bc))
                    ass[(ar, ac)] = set([change]).union(rm)
                else:
                    ass[(ar, ac)] = set()
                    break
            elif grid[ar][ac] == "#":
                ass[(ar, ac)] = set()
                break
            else:
                change = ((br, bc), (br + dr, bc))
                ass[(ar, ac)] = set([change])
        if all(len(s) != 0 for s in ass.values()):
            return set().union(*ass.values())
        else:
            return set()


def part2(intext: str) -> int:
    grid, moves = parse_grid_and_moves(intext)
    grid = resize_grid(grid)
    boxes = find_boxes(grid)
    rr, rc = find_robot(grid)

    for move in moves:
        # print_grid(grid)
        # print(move)

        dr, dc = move_map[move]
        nr, nc = (rr + dr, rc + dc)

        if grid[nr][nc] == "#":
            continue
        elif grid[nr][nc] == ".":
            grid[rr][rc] = "."
            grid[nr][nc] = "@"
            rr, rc = nr, nc
        else:
            er, ec = (nr, nc) if grid[nr][nc] == "[" else (nr, nc - 1)
            mss = move_box((er, ec), (dr, dc), grid, boxes)
            if mss:
                ms = {o: n for o, n in mss}
                new_boxes = set()
                for b in boxes:
                    new_boxes.add(ms[b] if b in ms else b)
                boxes = new_boxes
                for r in range(len(grid)):
                    for c in range(len(grid[0])):
                        if grid[r][c] in "[]":
                            grid[r][c] = "."
                for br, bc in boxes:
                    grid[br][bc] = "["
                    grid[br][bc + 1] = "]"
                grid[rr][rc] = "."
                grid[nr][nc] = "@"
                rr, rc = nr, nc

    total = 0
    for br, bc in boxes:
        total += (100 * br) + bc
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


# wow this code is ugly
