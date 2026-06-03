import AOC2021Lean.Solution

namespace AOC2021Lean.Days.Day01

def parseInput (input : String) : List Int :=
  input.splitOn "\n"
  |>.filterMap String.toInt?

def part1 (input : List Int) : Int :=
  part1Aux input.tail input.head! 0
  where
    part1Aux : List Int → Int → Int → Int
    | [], _, count => count
    | x :: xs, prev, count =>
        let count := if x > prev then count + 1 else count
        part1Aux xs x count

def part2 (input : List Int) : Int :=
  part2Aux (input.drop 3) input.head! (input.getD 1 0) (input.getD 2 0) 0
  where
    part2Aux : List Int → Int → Int → Int → Int → Int
    | [], _, _, _, count => count
    | x :: xs, prev1, prev2, prev3, count =>
        let count := if x > prev1 then count + 1 else count
        part2Aux xs prev2 prev3 x count

def solve (input : String) : String × String :=
  let nums := parseInput input
  let sol1 := part1 nums
  let sol2 := part2 nums
  (toString sol1, toString sol2)

def solution : AOC2021Lean.Solution :=
{
  solve := solve
}

end AOC2021Lean.Days.Day01
