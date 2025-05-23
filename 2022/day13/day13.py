from functools import cmp_to_key
from typing import Any


def get_pairs(filepath: str):
    with open(filepath) as f:
        data = f.read()
    pairs = data.strip().split('\n\n')
    return pairs


def get_packets(filepath: str):
    with open(filepath) as f:
        data = f.read()
    data = '\n'.join(data.strip().split('\n\n'))
    packets = data.split('\n')
    return packets


def compare_int(a: int, b: int) -> bool | None:
    # print(f'int compare: {a} with {b}')
    if a < b:
        return True
    elif a > b:
        return False
    else:
        return None


def compare_list(list_a: list[Any], list_b: list[Any]) -> bool | None:
    # print(f'list compare: {list_a} with {list_b}')
    while list_a and list_b:
        elem_a, elem_b = list_a.pop(0), list_b.pop(0)
        # print(f'elem compare: {elem_a} with {elem_b}')
        match elem_a, elem_b:
            case int(), int():
                result = compare_int(elem_a, elem_b)
                if result is not None:
                    return result
                else:
                    continue
            case int(), list():
                elem_a = [elem_a]
                result = compare_list(elem_a, elem_b)
                if result is not None:
                    return result
                else:
                    continue
            case list(), int():
                elem_b = [elem_b]
                result = compare_list(elem_a, elem_b)
                if result is not None:
                    return result
                else:
                    continue
            case list(), list():
                result = compare_list(elem_a, elem_b)
                if result is not None:
                    return result
                else:
                    continue
            case _:
                raise Exception('wtf')
    # print(len(list_a), len(list_b))
    if len(list_a) == 0 and len(list_b) > 0:
        return True
    elif len(list_a) > 0 and len(list_b) == 0:
        return False
    else:
        return None


def wrap_comp(sa: str, sb: str) -> int:
    packet_a = eval(sa)
    packet_b = eval(sb)
    result = compare_list(packet_a, packet_b)
    match result:
        case True:
            return -1
        case None:
            return 0
        case False:
            return 1


def validate(pair: str) -> bool | None:
    sa, sb = pair.split('\n')
    packet_a = eval(sa)
    packet_b = eval(sb)
    return compare_list(packet_a, packet_b)


if __name__ == '__main__':
    pairs = get_pairs('day13/input.txt')

    # part1
    valid_packet_indeces: list[int] = []
    for i, pair in enumerate(pairs, start=1):
        b = validate(pair)
        assert b is not None
        if b == True:
            valid_packet_indeces.append(i)
        print(pair, '\n', b, '\n')
    print(sum(valid_packet_indeces))

    # part2
    packets = get_packets('day13/input.txt')
    packets.append('[[2]]')
    packets.append('[[6]]')
    sp = sorted(packets, key=cmp_to_key(wrap_comp))
    print(*sp, sep='\n')
    di: list[int] = []
    for i, p in enumerate(sp, start=1):
        if p in ['[[2]]', '[[6]]']:
            di.append(i)
    print(di)
    print(di[0] * di[1])


"""
a little bit messy, but i got there kind of easily
"""
