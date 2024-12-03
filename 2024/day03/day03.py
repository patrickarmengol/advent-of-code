import sys
import re


def part1(intext: str) -> int:
    total = 0
    for a, b in re.findall(r"mul\((\d{1,3}),(\d{1,3})\)", intext):
        total += int(a) * int(b)
    return total


def part2(intext: str) -> int:
    enabled = True
    total = 0
    for a, b, dont, do in re.findall(
        r"mul\((\d{1,3}),(\d{1,3})\)|(don't\(\))|(do\(\))", intext
    ):
        if dont:
            enabled = False
        elif do:
            enabled = True
        elif enabled:
            total += int(a) * int(b)

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
