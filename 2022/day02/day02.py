move_points = {
    'X': 1,
    'Y': 2,
    'Z': 3,
    'A': 1,
    'B': 2,
    'C': 3
}


def p1outcome(opp_move: str, your_move: str) -> int:
    opp_val: int = move_points[opp_move]
    your_val: int = move_points[your_move]
    diff: int = opp_val - your_val
    if diff in (1, -2):
        return your_val + 0
    elif diff == 0:
        return your_val + 3
    else:
        return your_val + 6


def part1() -> None:
    score: int = 0
    with open('day02/input.txt') as f:
        rounds: list[str] = f.readlines()
        for round in rounds:
            round = round.rstrip('\n')
            score += p1outcome(*round.split())
    print(score)


def p2outcome(opp_move: str, ldw: str) -> int:
    move_arr: str = 'ABC'
    ldw_map: dict[str, int] = {'X': -1, 'Y': 0, 'Z': -2}
    your_move = move_arr[(move_points[opp_move] - 1) + ldw_map[ldw]]
    return p1outcome(opp_move, your_move)


def part2() -> None:
    score: int = 0
    with open('day02/input.txt') as f:
        rounds: list[str] = f.readlines()
        for round in rounds:
            round = round.rstrip('\n')
            score += p2outcome(*round.split())
    print(score)


if __name__ == '__main__':
    part1()
    part2()


"""
my initial answer was just to manually map pairs of letters to outcome values
i mean there are only 9 combinations and mentally calculating the result score isn't bad
but then i was like, let's do this "properly"
that's how i made this mess of a solution
"""
