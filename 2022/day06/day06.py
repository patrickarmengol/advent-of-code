def process_input(inpath: str) -> str:
    with open(inpath) as f:
        datastream_buf = f.read().rstrip('\n')
        return datastream_buf


def part1(data: str) -> None:
    for i in range(3, len(data)):
        if len(set(data[i-3:i+1])) == 4:
            print(i + 1, data[i-3:i+1])
            break


def part2(data: str) -> None:
    for i in range(13, len(data)):
        if len(set(data[i-13:i+1])) == 14:
            print(i + 1, data[i-13:i+1])
            break


if __name__ == '__main__':
    data = process_input('day06/input.txt')
    part1(data)
    part2(data)
