def process_input(filepath: str) -> list[list[tuple[int, int]]]:
    with open(filepath) as f:
        data = f.read()
    lines = data.strip().split('\n')
    walls: list[list[tuple[int, int]]] = [[eval(f'({point})') for point in line.split(' -> ')] for line in lines]
    return walls


def create_grid(walls: list[list[tuple[int, int]]]) -> list[list[str]]:
    floor = 0
    rightbound = 0
    for wall in walls:
        for c, r in wall:
            if r > floor:
                floor = r
            if c > rightbound:
                rightbound = c
    floor += 2
    rightbound += 5
    grid = [['.' if r < floor else '#' for _c in range(rightbound*2)] for r in range(floor+1)]
    for wall in walls:
        for i in range(1, len(wall)):
            point_a, point_b = wall[i-1], wall[i]
            min_c = min(point_a[0], point_b[0])
            max_c = max(point_a[0], point_b[0])
            for c in range(min_c, max_c+1):
                grid[point_a[1]][c] = '#'
            min_r = min(point_a[1], point_b[1])
            max_r = max(point_a[1], point_b[1])
            for r in range(min_r, max_r):
                grid[r][point_a[0]] = '#'
    return grid


def drop_sand(grid: list[list[str]]) -> int:
    count = 0
    grain: tuple[int, int] = (500, 0)
    # while grain[1] < len(grid)-1:
    while True:
        grain_c, grain_r = grain
        # check down
        if grid[grain_r + 1][grain_c] == '.':
            grain = (grain_c, grain_r + 1)
        # check downleft
        elif grid[grain_r + 1][grain_c - 1] == '.':
            grain = (grain_c - 1, grain_r + 1)
        # check downright
        elif grid[grain_r + 1][grain_c + 1] == '.':
            grain = (grain_c + 1, grain_r + 1)
        # rest
        else:
            count += 1
            grid[grain_r][grain_c] = 'o'
            if (grain_c, grain_r) == (500, 0):
                break
            grain = (500, 0)

    return count


def grid_print(grid: list[list[str]], left_bound: int, right_bound: int) -> None:
    for line in grid:
        print(*line[left_bound: right_bound], sep='')


if __name__ == '__main__':
    walls = process_input('day14/input.txt')
    grid = create_grid(walls)
    grid_print(grid, 400, 600)
    rounds = drop_sand(grid)
    grid_print(grid, 400, 600)
    print(rounds)


"""
i ended up breaking the logic for part1 but whatevs
falling sand sim is a classic
i could definitely do this more cleanly
"""
