import sys
import re


def parse_data(intext: str) -> tuple[tuple[int, int, int], list[int]]:
    regtxt, optxt = intext.split("\n\n")
    a, b, c = re.findall(r"\d+", regtxt)
    a, b, c = int(a), int(b), int(c)
    prog = re.findall(r"(\d)", optxt)
    prog = [int(x) for x in prog]
    return ((a, b, c), prog)


def part1(intext: str) -> str:
    (a, b, c), prog = parse_data(intext)

    def combo(n: int) -> int:
        match n:
            case 0 | 1 | 2 | 3:
                return n
            case 4:
                return a
            case 5:
                return b
            case 6:
                return c
            case 7:
                assert False
            case _:
                assert False

    i = 0
    out_buf = []
    while i < len(prog) - 1:
        op, oa = prog[i], prog[i + 1]
        # print(op, oa)
        match op:
            case 0:  # adv
                print(f"{a} // (2 ** {combo(oa)})")
                a = a // (2 ** combo(oa))
                print(a)
            case 1:  # bxl
                b = b ^ oa
            case 2:  # bst
                b = combo(oa) % 8
            case 3:  # jnz
                if a != 0:
                    i = oa // 2
                    continue
            case 4:  # bxc
                b = b ^ c
            case 5:  # out
                out_buf.append(combo(oa) % 8)
            case 6:  # bdv
                b = a // (2 ** combo(oa))
            case 7:  # cdv
                c = a // (2 ** combo(oa))
        i += 2

    return ",".join(str(x) for x in out_buf)


def part2(intext: str) -> int:
    # 2,4,1,7,7,5,0,3,4,4,1,7,5,5,3,0

    def run(prog: list[int], a: int) -> list[int]:
        b, c = 0, 0
        out: list[int] = []

        def combo(n: int) -> int:
            match n:
                case 0 | 1 | 2 | 3:
                    return n
                case 4:
                    return a
                case 5:
                    return b
                case 6:
                    return c
                case 7:
                    assert False
                case _:
                    assert False

        i = 0
        while 0 <= i < len(prog) - 1:
            op, oa = prog[i], prog[i + 1]
            match op:
                case 0:  # adv : a // 2^combo(oa) -> a
                    a = a // (2 ** combo(oa))
                case 1:  # bxl : b ^ oa -> b
                    b = b ^ oa
                case 2:  # bst : (combo(oa) % 8) -> b
                    b = combo(oa) % 8
                case 3:  # jnz : if a 0, noop; if a, goto oa
                    if a != 0:
                        i = oa // 2
                        continue
                case 4:  # bxc : b ^ c -> b
                    b = b ^ c
                case 5:  # out : combo(oa) % 8 -> print
                    # print(f"outputting {combo(oa) % 8}")
                    out.append(combo(oa) % 8)
                case 6:  # bdv : a // 2^combo(oa) -> b
                    b = a // (2 ** combo(oa))
                case 7:  # cdv : a // 2^combo(oa) -> c
                    c = a // (2 ** combo(oa))
            i += 2
        return out

    _, prog = parse_data(intext)

    target = prog[::-1]

    def find_a(a: int, d: int) -> int:
        if d == len(target):
            return a
        for i in range(8):
            cand = a * 8 + i
            output = run(prog, cand)
            if output and output[0] == target[d]:
                res = find_a(cand, d + 1)
                if res != 0:
                    return res
        return 0

    return find_a(0, 0)


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
