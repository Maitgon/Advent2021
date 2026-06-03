import AOC2021Lean
import AOC2021Lean.Days.Day01
import Init.Data.String.Basic

open AOC2021Lean

def getSolution? : Nat → Option Solution
| 1 => some AOC2021Lean.Days.Day01.solution
| _ => none

def formatDay (day : Nat) : String :=
  let s := toString day
  if s.length == 1 then
    "0" ++ s
  else
    s

def runDay (day : Nat) : IO Unit := do
  match getSolution? day with
  | none =>
      IO.println s!"Day {day} not implemented"

  | some solution =>


        let input ← IO.FS.readFile s!"AOC2021Lean/Input/day{formatDay day}.txt"

        let startTime ← IO.monoNanosNow
        let (p1, p2) := solution.solve input
        let endTime ← IO.monoNanosNow

        let ns := endTime - startTime

        IO.println s!"Part 1: {p1}"
        IO.println s!"Part 2: {p2}"
        IO.println s!"Time: ({ns / 1000} µs)"

def runRange (startDay endDay : Nat) : IO Unit := do
  for day in [startDay : endDay + 1] do
    runDay day
    IO.println ""

def parseNat? (s : String) : Option Nat :=
  s.toNat?

def main (args : List String) : IO Unit := do

  match args with

  | [dayStr] =>
      match parseNat? dayStr with
      | none =>
          IO.println "Invalid day"
      | some day =>
          runDay day

  | [startStr, endStr] =>
      match parseNat? startStr, parseNat? endStr with
      | some startDay, some endDay =>
          runRange startDay endDay
      | _, _ =>
          IO.println "Invalid range"

  | _ =>
      IO.println "Usage:"
      IO.println "  lake run 5"
      IO.println "  lake run 1 10"
