

def lines(filepath: str) -> list[str]:
    with open(filepath) as f:
        return f.read().rstrip('\n').split('\n')
