import sys


def parse_layout(intext: str) -> tuple[list[int], list[int]]:
    files = list()
    frees = list()
    i = 0
    for i in range(len(intext)):
        if i % 2 == 0:
            files.append(int(intext[i]))
        else:
            frees.append(int(intext[i]))
    return files, frees


def part1(intext: str) -> int:
    files, frees = parse_layout(intext)

    disk = []

    i = 0
    while i < len(files):
        disk.extend([i] * files[i])

        if i < len(files) - 1:
            b = len(files) - 1
            while frees[i] > 0:
                while files[b] == 0:
                    files.pop()
                    b -= 1
                frees[i] -= 1
                files[b] -= 1
                disk.append(b)

        i += 1

    return sum(i * disk[i] for i in range(len(disk)))


def part2(intext: str) -> int:
    files, frees = parse_layout(intext)

    disk: list[int] = []

    fills: list[list[tuple[int, int]]] = [[] for _ in range(len(frees))]

    for i in range(len(files) - 1, -1, -1):
        for j in range(0, i):
            if frees[j] >= files[i]:
                fills[j].append((i, files[i]))
                frees[j] -= files[i]  # reduce free space in filled spot
                frees[i - 1] += files[i]  # extend free space to left of moved
                files[i] = 0  # zero out moved

    for i in range(0, len(files)):
        disk.extend([i] * files[i])

        if i < len(frees):
            for fill in fills[i]:
                disk.extend([fill[0]] * fill[1])
            disk.extend([0] * frees[i])

    return sum(i * disk[i] for i in range(len(disk)))


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


# performance issues likely due to utilization of disk.extend([x] * y)
