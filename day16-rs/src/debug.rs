use aoc::GridExt;

use crate::pathstruct::Path;

pub fn debug_grid(mut grid: aoc::Grid<char>, paths: Vec<Path>) {
    for path in paths {
        *grid.at_point_mut(*path.cells.last().unwrap()) = path.dir.as_direction_char();
    }

    for y in 0..grid.height() {
        for x in 0..grid.width() {
            print!("{}", grid.at(x, y));
        }
        println!();
    }
}
