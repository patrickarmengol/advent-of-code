# advent of code

Here are my solutions to the Advent of Code problems.

## history

Last year was my first attempt at the challenge. My run came to a halt at problem 16, which stumped me for the majority the day before I decided to go hiking instead. I was using Python and my code was quite messy and I kept breaking the solution for part1 in order to do part2.

I'm currently working on the problems for 2023 as they go live. I decided to use Go (Golang) this year, as I'm already decently familiar with Python for DSA. Go doesn't have nearly as many conveniences built-in as Python does, but it's fairly easy (and fun) to implement things yourself, especially now with generics.

## setup

I'm a bit troubled by how I want to set up the code for this project. My current plan is as follows.

I make the year a go module that has helper packages for use in all days. A lot of other people seem to prefer to make each day a main package with a main function to run each part and get the output directly. Seems simple enough. But I think I want to use go's testing features in order to make it more robust. Each day will be it's own package. `solution.go` will house the code for solving parts 1 and 2. `solution_test.go` will call the solution functions and check against expected answers. But we are only given expected answers for the examples. For the actual user-unique solution, I can just run the test against _"UNKNOWN"_ and let it fail to get the output. When I find the correct answer, I put it right into the test's expected data so that I can modify/improve the code and check that it still works. This is also useful because occassionally I would want to modify the helper functions and check if any of the solutions break across all days. You know... the whole reason testing exists.

I know others have made tools that do things like automate pulling in the example/problem input into files, generating code from templates for each day, and some AI shenangians. I might look into implementing some of these myself.

## usage notes

process:

1. copy the template `day00`
2. fill in `example.txt` and `input.txt`
3. fill in the expected answer for `TestPart1Example`
4. implement solution for `Part1`
5. `go test -v ./day??`
6. if `TestPart1Example` passed, check result for `TestPart1Actual` on site
7. fill in the expected answer for `TestPart1Actual`
8. repeat from step 3 for part 2
