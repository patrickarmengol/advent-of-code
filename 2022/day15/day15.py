import re


def process_input(filepath: str) -> list[tuple[tuple[int, int], tuple[int, int]]]:
    with open(filepath) as f:
        data = f.read()
    lines = data.strip().split('\n')
    sb_pairs: list[tuple[tuple[int, int], tuple[int, int]]] = []
    for line in lines:
        numbers = re.findall(r'-?\d+', line)
        sensor_x, sensor_y, beacon_x, beacon_y = [int(_) for _ in numbers]
        sb_pairs.append(((sensor_x, sensor_y), (beacon_x, beacon_y)))
    return sb_pairs


def man_dist(ax: int, ay: int, bx: int, by: int) -> int:
    return abs(ax - bx) + abs(ay - by)


def edges(sx: int, sy: int, smd: int) -> set[tuple[int, int]]:
    es: set[tuple[int, int]] = set()
    omd = smd + 1
    for dx in range(-omd, omd+1):
        x = sx + dx
        y = sy-(omd-abs(dx))
        es.add((x, y))
        y = sy+(omd-abs(dx))
        es.add((x, y))
    return es


def part1(sb_pairs: list[tuple[tuple[int, int], tuple[int, int]]], row_index: int) -> int:
    sdict: dict[tuple[int, int], int] = dict()
    beaks: set[tuple[int, int]] = set()
    for (sx, sy), (bx, by) in sb_pairs:
        md = man_dist(sx, sy, bx, by)
        sdict[(sx, sy)] = md
        beaks.add((bx, by))
    inner_beaks = sum([by == row_index for bx, by in beaks])

    count_set: set[int] = set()
    for (sx, sy), smd in sdict.items():
        dif_y = abs(row_index - sy)
        if dif_y <= smd:
            minx = sx - (smd - dif_y)
            maxx = sx + (smd - dif_y)
            print(minx, maxx)
            for i in range(minx, maxx+1):
                count_set.add(i)
    return len(count_set) - inner_beaks


def part2(sb_pairs: list[tuple[tuple[int, int], tuple[int, int]]], leftx: int, rightx: int, upy: int, downy: int) -> int | None:
    sdict: dict[tuple[int, int], int] = dict()

    for (sx, sy), (bx, by) in sb_pairs:
        md = man_dist(sx, sy, bx, by)
        sdict[(sx, sy)] = md

    for (sx, sy), smd in sdict.items():
        #print(f'\n edges for {sx},{sy} with {smd = }')
        es = edges(sx, sy, smd)
        for ex, ey in es:
            if leftx <= ex <= rightx and upy <= ey <= downy:
                #print(ex, ey)
                for (ssx, ssy), ssmd in sdict.items():
                    if man_dist(ex, ey, ssx, ssy) <= ssmd:
                        break
                else:
                    print(ex, ey)
                    return ex * 4000000 + ey
    return None


if __name__ == '__main__':

    # # part 1
    # sb_pairs = process_input('day15/example.txt')
    # print(part1(sb_pairs, 10))

    # sb_pairs = process_input('day15/input.txt')
    # print(part0(sb_pairs, 2000000))

    # part2
    # sb_pairs = process_input('day15/example.txt')
    # print(part2(sb_pairs, 0, 20, 0, 20))

    sb_pairs = process_input('day15/input.txt')
    print(part2(sb_pairs, 0, 4000000, 0, 4000000))


"""
i got stuck on part 1 because i was parsing the input incorrectly (forgot the minus sign in regex pattern)
part 2 really stumped me for a while though
i was like, well iterating 4_000_000 ** 2 items is definitely not going to be the answer but wtf else do i do?
but then i though about how to narrow down which coords to test, 
and since there can only be 1 valid location,it would have to be right outside the range of one of the sources
"""
