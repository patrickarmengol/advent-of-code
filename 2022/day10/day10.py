
with open('day10/input.txt') as f:
    data: str = f.read()
    lines = data.rstrip('\n').split('\n')

register = 1  # sprite pos

cycles = {
    'noop': 1,
    'addx': 2,
}


cur_ins = None
cycles_left = 0
total_ss = 0
i = 0
cycle_index = 0
while i < len(lines):
    cycle_index += 1
    instruction = lines[i]
    crt_index = cycle_index % 40 - 1
    print('#' if abs(crt_index - register) <= 1 else '.', end=('\n' if cycle_index % 40 == 0 else ''))
    if (cycle_index - 20) % 40 == 0:
        total_ss += cycle_index * register
        #print(f'{cycle_index = }\t{register = }\t->\t{cycle_index * register}')
    if not cur_ins:
        operation = instruction.split()[0]
        cur_ins = instruction
        cycles_left = cycles[operation]
        #print(f'starting: {cur_ins} @ cycle {cycle_index}; register = {register}')
    cycles_left -= 1
    if cycles_left == 0:
        match cur_ins.split()[0]:
            case 'noop':
                pass
            case 'addx':
                register += int(cur_ins.split()[1])
            case _:
                print('wtf')
        #print(f'ending {cur_ins} @ cycle {cycle_index}; register = {register}')
        # print()
        cur_ins = None
        i += 1

print(total_ss)


"""
###..#..#..##..#..#.#..#.###..####.#..##
#..#.#..#.#..#.#.#..#..#.#..#.#....#.#..
#..#.#..#.#..#.##...####.###..###..##..#
###..#..#.####.#.#..#..#.#..#.#....#.#.#
#.#..#..#.#..#.#.#..#..#.#..#.#....#.#..
#..#..##..#..#.#..#.#..#.###..####.#..#.

kind of easy day
"""
