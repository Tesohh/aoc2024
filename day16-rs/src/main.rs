pub mod debug;
pub mod findpath;
pub mod pathstruct;

fn main() {
    let grid = aoc::Input::from_args().char_grid();
    findpath::find_path(grid);
}
