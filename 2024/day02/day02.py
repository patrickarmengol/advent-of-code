import sys


def parse_reports(intext: str) -> list[list[int]]:
    return [[int(x) for x in line.split()] for line in intext.split("\n")]


def part1(intext: str) -> int:
    reports = parse_reports(intext)

    count = 0

    for report in reports:
        # check inc/dec
        if not (sorted(report) == report or sorted(report, reverse=True) == report):
            continue

        # check steps not too big/small
        for i in range(1, len(report)):
            diff = abs(report[i] - report[i - 1])
            if not (1 <= diff <= 3):
                break
        else:
            count += 1

    return count


def part2(intext: str) -> int:
    reports = parse_reports(intext)

    count = 0

    for orig_report in reports:
        # from each original report, create set of alternative reports, each missing single level
        for i in range(0, len(orig_report)):
            report = orig_report[:i] + orig_report[i + 1 :]

            # check inc/dec
            if not (sorted(report) == report or sorted(report, reverse=True) == report):
                continue

            # check steps not too big/small
            for i in range(1, len(report)):
                diff = abs(report[i] - report[i - 1])
                if not (1 <= diff <= 3):
                    break
            else:
                count += 1
                break  # break here if found one valid in set of alternative reports

    return count


# read input and call part functions
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
