from __future__ import annotations


class TreeNode:
    def __init__(self, name: str, kind: str, size: int = 0, parent: TreeNode | None = None):
        self.name = name
        self.kind = kind
        self.size = size
        self.parent = parent
        self.children: list[TreeNode] = []


def process_input(filepath: str) -> tuple[str, list[str]]:
    with open(filepath) as f:
        data = f.read()
        lines = data.rstrip('\n').split('\n')
    return data, lines


def build_tree(lines: list[str]) -> TreeNode:
    filesystem_root: TreeNode = TreeNode('/', kind='dir')
    cwd: TreeNode = filesystem_root
    for line in lines:
        if line.startswith('$ cd'):
            desired = line.split()[-1]
            if desired == '/':
                continue
            elif desired == '..':
                if cwd.parent:
                    cwd = cwd.parent
                else:
                    raise Exception('wtf')
            else:
                for child in cwd.children:
                    if child.name == desired:
                        cwd = child
                        break
                else:
                    new_child = TreeNode(desired, kind='dir', parent=cwd)
                    cwd.children.append(new_child)
                    cwd = new_child
        elif line.startswith('dir'):
            child_dir_name = line.split()[-1]
            if child_dir_name not in [child.name for child in cwd.children]:
                new_child = TreeNode(child_dir_name, kind='dir', parent=cwd)
                cwd.children.append(new_child)
        elif line.split()[0].isnumeric():
            child_file_size, child_file_name = line.split()
            if child_file_name not in [child.name for child in cwd.children]:
                new_child = TreeNode(child_file_name, kind='file', size=int(child_file_size), parent=cwd)
                cwd.children.append(new_child)
    return filesystem_root


def walk_size(node: TreeNode, dir_sizes: list[tuple[str, int]]) -> int:
    if len(node.children) > 0:
        node.size = sum([walk_size(child, dir_sizes) for child in node.children])
        dir_sizes.append((node.name, node.size))
    return node.size


def walk_print(node: TreeNode, leader: str = '') -> None:
    print(f'{leader}- {node.name} ({node.kind}{(", size=" + str(node.size)) if node.kind in ["file", "dir"] else ""})')
    if len(node.children) > 0:
        [walk_print(child, leader=leader+'  ') for child in node.children]


if __name__ == '__main__':
    data, lines = process_input('day07/input.txt')
    filesystem = build_tree(lines)
    sizes: list[tuple[str, int]] = []
    walk_size(filesystem, sizes)
    walk_print(filesystem)

    part1_sum = 0
    for dirname, dirsize in sizes:
        if dirsize <= 100000:
            print(dirname, dirsize)
            part1_sum += dirsize
    print(f'part1: {part1_sum}')

    disk_cap = 70000000
    update_req = 30000000
    free_space = disk_cap - filesystem.size
    print(f'free_space: {free_space}')
    needed_deletion = update_req - free_space
    print(f'needed_deletion: {needed_deletion}')
    dir_to_delete = min([size for size in sizes if size[1] >= needed_deletion], key=lambda x: x[1])
    print(f'part2: {dir_to_delete}')


"""
holy heck did i complicate things
i'm still kind of getting the hang of type hints

found a crazy sol on reddit:

stack = []
sizes = []

def up():
    sizes.append(stack.pop(-1))
    if stack:
        stack[-1] += sizes[-1]

for line in open("../inputs/07.txt").readlines():
    match line.strip().split():
        case "$", "cd", "..": up()
        case "$", "cd", _: stack.append(0)
        case "$", _: pass
        case "dir", _: pass
        case size, _: stack[-1] += int(size)

while stack:
    up()

print(sum(s for s in sizes if s <= 100000))
print(min(s for s in sizes if s >= (max(sizes) - 40000000)))
"""
