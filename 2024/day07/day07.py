from itertools import zip_longest, chain
import sys


def parse_equations(intext: str) -> list[tuple[int, list[int]]]:
    eqs = []
    for line in intext.split("\n"):
        res, nums = line.split(": ")
        res = int(res)
        nums = [int(x) for x in nums.split(" ")]
        eqs.append((res, nums))
    return eqs


def part1(intext: str) -> int:
    eqs = parse_equations(intext)

    def back(cur: int, goal: int, nums: list[int], ops: str, i: int) -> bool:
        if i == len(nums) - 1 and cur == goal:
            return True
        if cur > goal or i == len(nums) - 1:
            return False

        return back(cur + nums[i + 1], goal, nums, ops + "+", i + 1) or back(
            cur * nums[i + 1], goal, nums, ops + "*", i + 1
        )

    total = 0
    for res, nums in eqs:
        if back(nums[0], res, nums, "", 0):
            total += res

    return total


def part2(intext: str) -> int:
    eqs = parse_equations(intext)

    def back(cur: int, goal: int, nums: list[int], ops: str, i: int) -> bool:
        if i == len(nums) - 1 and cur == goal:
            # print(
            #     f"{goal}: {' '.join(f'{n} {o if o else ""}' for (n, o) in zip_longest(nums, ops))}"
            # )
            return True
        if cur > goal or i == len(nums) - 1:
            return False

        return (
            back(cur + nums[i + 1], goal, nums, ops + "+", i + 1)
            or back(cur * nums[i + 1], goal, nums, ops + "*", i + 1)
            or back(int(str(cur) + str(nums[i + 1])), goal, nums, ops + "|", i + 1)
        )

    total = 0
    for res, nums in eqs:
        if back(nums[0], res, nums, "", 0):
            total += res

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
