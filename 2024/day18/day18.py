from queue import PriorityQueue
import sys
import re
from typing import Callable

sys.setrecursionlimit(9999)  # lol


def parse_bytes(intext: str) -> list[tuple[int, int]]:
    bs = re.findall(r"(\d+),(\d+)", intext)
    return [(int(y), int(x)) for (x, y) in bs]  # r,c order


def neighbors(
    point: tuple[int, int], size: int, cond: Callable
) -> list[tuple[int, int]]:
    ns = []
    (pr, pc) = point
    for dr, dc in [(0, -1), (0, 1), (-1, 0), (1, 0)]:
        nr, nc = (pr + dr, pc + dc)
        if 0 <= nr < size and 0 <= nc < size and cond((nr, nc)):
            ns.append((nr, nc))
    return ns


def heuristic(a: tuple[int, int], b: tuple[int, int]) -> int:
    return abs(a[0] - b[0]) + abs(a[1] - b[1])


def part1(intext: str) -> int:
    size = 7 if sys.argv[2] == "example.txt" else 71
    steps = 12 if sys.argv[2] == "example.txt" else 1024

    bs = parse_bytes(intext)
    dropped = set(bs[:steps])

    # a*
    start = (0, 0)
    goal = (size - 1, size - 1)

    frontier = PriorityQueue()
    frontier.put((0, start))

    came_from = dict()
    cost_so_far = dict()
    came_from[start] = None
    cost_so_far[start] = 0

    while not frontier.empty():
        cur = frontier.get()[1]

        if cur == goal:
            break

        for n in neighbors(cur, size, lambda x: x not in dropped):
            new_cost = cost_so_far[cur] + 1
            if n not in cost_so_far or new_cost < cost_so_far[n]:
                cost_so_far[n] = new_cost
                priority = new_cost + heuristic(goal, n)
                frontier.put((priority, n))
                came_from[n] = cur

    return cost_so_far[goal]


def part2(intext: str) -> str:
    size = 7 if sys.argv[2] == "example.txt" else 71

    bs = parse_bytes(intext)

    def can_reach(
        start: tuple[int, int], goal: tuple[int, int], obs: set[tuple[int, int]]
    ) -> bool:
        visited: set[tuple[int, int]] = set()

        def dfs(coord: tuple[int, int], target: tuple[int, int]) -> bool:
            if coord == target:
                return True
            visited.add(coord)
            return any(
                dfs((nr, nc), target)
                for (nr, nc) in neighbors(coord, size, lambda x: x not in obs)
                if (nr, nc) not in visited
            )

        return dfs(start, goal)

    start = (0, 0)
    goal = (size - 1, size - 1)

    lo = 0
    hi = len(bs)

    while lo < hi:
        mid = (lo + hi) // 2
        if can_reach(start, goal, set(bs[:mid])):
            lo = mid + 1
        else:
            hi = mid

    rr, rc = bs[lo - 1]
    return f"{rc},{rr}"


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
