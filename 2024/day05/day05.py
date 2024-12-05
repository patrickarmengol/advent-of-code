from functools import cmp_to_key
import sys
from collections import defaultdict


def parse_rules_updates(
    intext: str,
) -> tuple[defaultdict[str, set[str]], list[list[str]]]:
    ruletext, updatetext = intext.split("\n\n")

    rules: defaultdict[str, set[str]] = defaultdict(set[str])
    for rule in ruletext.split("\n"):
        a, b = rule.split("|")
        rules[a].add(b)

    updates = []
    for update in updatetext.split("\n"):
        updates.append(update.split(","))

    return rules, updates


def part1(intext: str) -> int:
    rules, updates = parse_rules_updates(intext)

    mids = []
    for update in updates:
        page_indeces = {update[i]: i for i in range(len(update))}
        valid = True
        for page in page_indeces:
            if page in rules and any(
                page_indeces[page] > page_indeces[rule]
                for rule in rules[page]
                if rule in page_indeces
            ):
                valid = False
                break
        if valid:
            mids.append(update[len(update) // 2])

    return sum(int(x) for x in mids)


def part2(intext: str) -> int:
    rules, updates = parse_rules_updates(intext)

    def compare_pages(a: str, b: str) -> int:
        if a == b:
            return 0
        elif b in rules[a]:
            return -1
        else:
            return 1

    mids = []
    for update in updates:
        page_indeces = {update[i]: i for i in range(len(update))}
        for page in page_indeces:
            if page in rules and any(
                page_indeces[page] > page_indeces[rule]
                for rule in rules[page]
                if rule in page_indeces
            ):
                update = sorted(update, key=cmp_to_key(compare_pages))
                mids.append(update[len(update) // 2])
                break
    return sum(int(x) for x in mids)


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
