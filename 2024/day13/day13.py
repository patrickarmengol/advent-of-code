import sys
import re


def parse_machines(
    intext: str,
) -> list[tuple[tuple[int, int], tuple[int, int], tuple[int, int]]]:
    # returns [((ax1, ay1), (bx1, by1), (px1, py1)), ...]
    machines = []
    for machinetext in intext.split("\n\n"):
        lines = machinetext.split("\n")
        assert len(lines) == 3
        at, bt, pt = lines
        ax, ay = re.findall(r"\d+", at)
        bx, by = re.findall(r"\d+", bt)
        px, py = re.findall(r"\d+", pt)
        machines.append(((int(ax), int(ay)), (int(bx), int(by)), (int(px), int(py))))
    return machines


def cost(pressed: tuple[int, int]) -> int:
    return (3 * pressed[0]) + (1 * pressed[1])


def part1(intext: str) -> int:
    machines = parse_machines(intext)

    total = 0
    for (ax, ay), (bx, by), (px, py) in machines:
        tried = set()
        cache = {}

        def dfs(cur: tuple[int, int], path: tuple[int, int]):
            if path in tried:
                return
            tried.add(path)

            if cur == (px, py):
                cache[(px, py)] = (
                    min(cost(path), cache[(px, py)])
                    if (px, py) in cache
                    else cost(path)
                )
                return
            elif cur[0] > px or cur[1] > py:
                return
            elif cur in cache:
                if cache[cur] > cost(cur):
                    cache[cur] = cost(cur)
                else:
                    return
            dfs((cur[0] + ax, cur[1] + ay), (path[0] + 1, path[1]))
            dfs((cur[0] + bx, cur[1] + by), (path[0], path[1] + 1))

        dfs((0, 0), (0, 0))
        total += cache[(px, py)] if (px, py) in cache else 0

    return total


def part2(intext: str) -> int:
    machines = parse_machines(intext)

    total = 0
    for (ax, ay), (bx, by), (px, py) in machines:
        px, py = px + 10000000000000, py + 10000000000000

        # system of equations
        # a * ax + b * bx = px
        # a * ay + b * by = py
        # ax, ay, bx, by, px, py are known
        # solve for a and b
        # cramer's rule
        # a = (px * by - py * bx) / (ax * by - ay * bx)
        # b = (ax * py - ay * px) / (ax * by - ay * bx)

        det = (ax * by) - (ay * bx)
        a = ((px * by) - (py * bx)) // det
        b = ((ax * py) - (ay * px)) // det
        if (a * ax + b * bx, a * ay + b * by) == (px, py):
            total += a * 3 + b

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
