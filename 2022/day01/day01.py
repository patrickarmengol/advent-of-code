

def main() -> None:
    with open('day01/input.txt') as f:
        elf_calories = [sum(int(meal) for meal in elf.split('\n'))
                        for elf in f.read().rstrip('\n').split('\n\n')]
        thiccest_elf = max(elf_calories)
        print(thiccest_elf)
        three_thiccest_elves = sorted(elf_calories, reverse=True)[:3]
        print(three_thiccest_elves)
        total_emergency_rations = sum(three_thiccest_elves)
        print(total_emergency_rations)


if __name__ == '__main__':
    main()
