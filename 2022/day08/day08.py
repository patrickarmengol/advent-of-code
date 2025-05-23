def process_input(filepath: str) -> list[list[int]]:
    with open(filepath) as f:
        data: str = f.read()
        lines = data.rstrip('\n').split('\n')
        grid: list[list[int]] = [[int(item) for item in line] for line in lines]
    return grid


def vis_matrix(grid: list[list[int]]) -> list[list[bool]]:
    def is_visible(row: int, col: int) -> bool:
        tree = grid[row][col]
        return all(tree > grid[row][x] for x in range(col+1, x_max+1)) or all(tree > grid[row][x] for x in range(col-1, -1, -1)) or all(
            tree > grid[y][col] for y in range(row+1, y_max+1)) or all(tree > grid[y][col] for y in range(row-1, -1, -1))

    y_max, x_max = len(grid)-1, len(grid[0])-1
    return [[is_visible(row, col) for col in range(x_max+1)] for row in range(y_max+1)]


def scenic_matrix(grid: list[list[int]]) -> list[list[int]]:
    def score(row: int, col: int) -> int:
        tree = grid[row][col]
        a = 0
        for t in [grid[row][x] for x in range(col+1, x_max+1)]:
            a += 1
            if t >= tree:
                break
        b = 0
        for t in [grid[row][x] for x in range(col-1, -1, -1)]:
            b += 1
            if t >= tree:
                break
        c = 0
        for t in [grid[y][col] for y in range(row+1, y_max+1)]:
            c += 1
            if t >= tree:
                break
        d = 0
        for t in [grid[y][col] for y in range(row-1, -1, -1)]:
            d += 1
            if t >= tree:
                break
        return a * b * c * d

    y_max, x_max = len(grid)-1, len(grid[0])-1
    return [[score(row, col) for col in range(1, x_max)] for row in range(1, y_max)]


if __name__ == '__main__':
    grid = process_input('day08/input.txt')
    print(*grid, sep='\n')
    vism = vis_matrix(grid)
    print(*vism, sep='\n')
    num_visible = sum([t for row in vism for t in row])
    print(f'{num_visible = }')
    scem = scenic_matrix(grid)
    print(*scem, sep='\n')
    prettiest = max(t for row in scem for t in row)
    print(f'{prettiest = }')


"""
yeah, i want to redo this one with numpy arrays
"""
