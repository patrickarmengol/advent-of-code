import numpy as np
import numpy.typing as npt

d = {
    'U': np.array([0, 1]),
    'D': np.array([0, -1]),
    'R': np.array([1, 0]),
    'L': np.array([-1, 0]),
}


def process_input(filepath: str) -> list[str]:
    with open(filepath) as f:
        data: str = f.read()
        lines = data.rstrip('\n').split('\n')
    return lines


def part1(lines: list[str]) -> None:
    h: npt.NDArray[np.int_] = np.array([0, 0])
    t: npt.NDArray[np.int_] = np.array([0, 0])
    visited = {str(np.array([0, 0]))}

    for line in lines:
        ins = line.split()
        direction, steps = ins[0], int(ins[1])
        delta: npt.NDArray[np.int_] = d[direction]
        for _ in range(steps):
            h += delta
            diff = h - t
            #print('move head', h, t, diff)
            if abs(diff[0]) > 1:
                t[1] = h[1]
                t[0] = h[0] - diff[0] // 2
            elif abs(diff[1]) > 1:
                t[0] = h[0]
                t[1] = h[1] - diff[1] // 2
            #new_diff = h - t
            #print('after tail', h, t, new_diff)
            # print()
            visited.add(str(t))
    #print(*visited, sep='\n')
    print(len(visited))


def part2(lines: list[str]) -> None:
    h: npt.NDArray[np.int_] = np.array([0, 0])
    ta: list[npt.NDArray[np.int_]] = [np.array([0, 0]) for _ in range(9)]
    visited = {str(np.array([0, 0]))}

    for line in lines:
        ins = line.split()
        direction, steps = ins[0], int(ins[1])
        delta: npt.NDArray[np.int_] = d[direction]
        for _ in range(steps):
            h += delta
            prev = h
            for t in ta:
                diff = prev - t
                if abs(diff[0]) > 1 and abs(diff[1]) > 1:
                    t[1] = prev[1] - diff[1] // 2
                    t[0] = prev[0] - diff[0] // 2
                elif abs(diff[0]) > 1:
                    t[1] = prev[1]
                    t[0] = prev[0] - diff[0] // 2
                elif abs(diff[1]) > 1:
                    t[0] = prev[0]
                    t[1] = prev[1] - diff[1] // 2
                else:
                    break
                prev = t
            else:
                visited.add(str(prev))
    print(len(visited))


if __name__ == '__main__':
    lines = process_input('day09/input.txt')
    part1(lines)
    part2(lines)


"""
what a mess, i decided to use numpy in order to more esily do vector operations, but i didn't even use them
idk what is going on with numpy typing, but pylance/pyright is yelling at me for weird reasons

i think an improvement to this would be to split the knot adjusting into its own function
"""
