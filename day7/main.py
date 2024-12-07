import dataclasses
import sys

@dataclasses.dataclass
class Equation:
    expected: int
    numbers: list[int]

    def calc(self, ops: list[str]) -> int:
        start = self.numbers[0]
        for i, op in enumerate(ops):
            if op == "+": start += self.numbers[i+1]
            if op == "*": start *= self.numbers[i+1]
            if op == "|": start = int(str(start) + str(self.numbers[i+1]))
        return start

    def is_correct(self, ops: list[str]) -> bool:
        return self.calc(ops) == self.expected

def get_ops(ops: str, len: int):
    if len==1: return ops
    r = []
    for oo in get_ops(ops,len-1):
        for o in ops: 
            r.append(oo+o)
    return r


equations: list[Equation] = []

with open(sys.argv[1]) as f:
    for line in f.readlines():
        parts = line.split(":")
        eq = Equation(int(parts[0]), [int(x.strip()) for x in parts[1].strip().split(" ")])
        equations.append(eq)


print(equations)

part1 = set()
for eq in equations:
    for ops in get_ops("+*|", len(eq.numbers)-1):
        if eq.is_correct(list(ops)):
            # print(f"found solution for {eq} with ops {ops}")
            part1.add((eq.expected, tuple(eq.numbers)))

sum = 0
for x in part1:
    sum += x[0]

print("part1",sum)
