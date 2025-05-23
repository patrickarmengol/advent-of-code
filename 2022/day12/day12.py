
def process_input(filepath: str) -> list[str]:
    with open(filepath) as f:
        data = f.read()
    lines = data.strip().split('\n')
    return lines


def get_neighbors(grid: list[list[str]], node: tuple[int, int], visited: set[tuple[int, int]]) -> list[tuple[int, int]]:
    neighbors: list[tuple[int, int]] = []
    for delta in [(-1, 0), (0, -1), (1, 0), (0, 1)]:
        px = node[0] + delta[0]
        py = node[1] + delta[1]
        if not (0 <= px < MAX_X and 0 <= py < MAX_Y):
            continue
        p = grid[px][py]
        n = grid[node[0]][node[1]]
        if tuple([px, py]) not in visited and (p == 'a' or (ord(p) - ord(n) <= 1 and p.islower()) or (n in 'yz' and p == 'E')):
            neighbors.append((px, py))
    return neighbors


def bfs(grid: list[list[str]], start: tuple[int, int]) -> list[tuple[int, int]]:
    marked = set([start])
    queue = [[start]]
    path: list[tuple[int, int]] = []

    while queue:
        path = queue.pop(0)
        node = path[-1]
        letter = grid[node[0]][node[1]]
        #print(f'{node = } {letter}')
        if letter == 'E':
            return path

        for neighbor in get_neighbors(grid, node, marked):
            #print(f'{neighbor = }')
            marked.add(neighbor)
            queue.append(path + [neighbor])
    raise Exception('no path')


def find_letter(grid: list[list[str]], letter: str) -> list[tuple[int, int]]:
    found: list[tuple[int, int]] = []
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if grid[i][j] == letter:
                found.append((i, j))
    return found


def pretty_path(path: list[tuple[int, int]]) -> None:
    print(path)
    print(len(path) - 1)
    pg = [[' ' for _ in line] for line in grid]
    for n in path:
        pg[n[0]][n[1]] = grid[n[0]][n[1]]
    for line in pg:
        print(''.join(line))


if __name__ == '__main__':
    lines = process_input('day12/input.txt')
    grid: list[list[str]] = [list(line) for line in lines]
    MAX_Y = len(grid[0])
    MAX_X = len(grid)
    print(*grid, sep='\n')

    # part 1
    path = bfs(grid, find_letter(grid, 'S')[0])
    # pretty_path(path)

    # part2
    min_path = [(i, i) for i in range(99999999)]
    for a in find_letter(grid, 'a'):
        try:
            print(f'start: {a}')
            path = bfs(grid, a)
            if len(path) < len(min_path):
                min_path = path
        except Exception as e:
            print(e)

    pretty_path(min_path)


"""
fun
"""
