from math import prod

with open('day11/input.txt') as f:
    data: str = f.read()


monkeys = {}
cur = 0
for section in data.rstrip('\n').split('\n\n'):
    for line in section.split('\n'):
        if line.startswith('Monkey'):
            cur = int(line.split()[1].rstrip(':'))
            monkeys[cur] = {'count': 0}
        else:
            print(line)
            k, v = line.strip().split(': ')
            match k:
                case 'Starting items':
                    monkeys[cur]['items'] = [int(x) for x in v.split(', ')]
                case 'Operation':
                    monkeys[cur]['operation'] = v.split('= ')[1]
                case 'Test':
                    monkeys[cur]['test'] = int(v.split()[-1])
                case 'If true':
                    monkeys[cur]['true'] = int(v.split()[-1])
                case 'If false':
                    monkeys[cur]['false'] = int(v.split()[-1])
                case _:
                    print('wtf')

print(monkeys)

part2_mod = prod([m['test'] for m in monkeys.values()])

for r in range(10000):
    # new round
    # print(r)
    for i in range(len(monkeys)):
        # new turn for each monkey
        m = monkeys[i]
        for item in m['items']:
            m['count'] += 1
            # inspect by doing operation
            old = item
            new = eval(m['operation'])
            #new = new // 3
            new = new % part2_mod
            if new % m['test'] == 0:
                monkeys[m['true']]['items'].append(new)
            else:
                monkeys[m['false']]['items'].append(new)
        m['items'] = []

print(monkeys)

counts = []
for id, m in monkeys.items():
    print(id, m['count'])
    counts.append(m['count'])

a, b = sorted(counts)[-2:]
print(f'{a * b = }')


"""
cool
"""
