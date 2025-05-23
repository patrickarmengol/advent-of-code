def process_input(filepath: str) -> tuple[list[list[str]], list[str]]:
    with open(filepath) as f:
        stacks_str, instructions_str = f.read().rstrip('\n').split('\n\n')

        stack_lines = stacks_str.split('\n')
        n_stacks = int(stack_lines[-1].split()[-1])
        stacks: list[list[str]] = [[] for _ in range(n_stacks)]
        for line in stack_lines[:-1]:
            for i in range(n_stacks):
                item = line[1 + i * 4]
                if item != ' ':
                    stacks[i].append(item)
        stacks = [stack[::-1] for stack in stacks]

        instructions = instructions_str.split('\n')

        return stacks, instructions


def part1(stacks: list[list[str]], instructions: list[str]) -> None:
    for instruction in instructions:
        amount, src, dst = [int(x)
                            for x in instruction.split() if x.isnumeric()]
        for _ in range(amount):
            stacks[dst-1].append(stacks[src-1].pop())
    print(''.join([stack[-1] for stack in stacks]))


def part2(stacks: list[list[str]], instructions: list[str]) -> None:
    print(stacks)
    for instruction in instructions:
        amount, src, dst = [int(x)
                            for x in instruction.split() if x.isnumeric()]
        crane_claw = [stacks[src-1].pop() for _ in range(amount)]
        crane_claw.reverse()
        stacks[dst-1].extend(crane_claw)
    print(''.join([stack[-1] for stack in stacks]))


if __name__ == '__main__':
    problem_input = 'day05/input.txt'
    part1(*process_input(problem_input))
    part2(*process_input(problem_input))
