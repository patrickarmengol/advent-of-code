# advent of code

here are my solutions to the Advent of Code problems

## usage notes

my current and not-so-optimal process:

1. copy the template `day00` to the new day; cd into the day
2. fill in `example.txt` and `input.txt`
3. fill in the expected answer for `TestPart1Example`
4. implement solution for `Part1`
5. `go test -v` or `go test -v -run Part1`
6. if `TestPart1Example` passed, check result for `TestPart1Actual` on site
7. fill in the expected answer for `TestPart1Actual`
8. repeat from step 3 for part 2
