pub mod hiker;
pub mod solve;

fn main() {
    let input = aoc::Input::from_args();
    let grid = input.i64_grid();

    dbg!(solve::solve(grid));
}
