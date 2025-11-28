use aoc::{GridExt, Vec2};

use crate::debug;
use crate::pathstruct::Path;

pub fn find_path(grid: aoc::Grid<char>) {
    let mut paths: Vec<Path> = vec![];

    for y in 0..grid.height() {
        for x in 0..grid.width() {
            if *grid.at(x, y) == 'S' {
                paths.push(Path {
                    cells: vec![Vec2 {
                        x: x as i64,
                        y: y as i64,
                    }],
                    dir: Vec2::RIGHT,
                    cost: 0,
                    retired: false,
                    finished: false,
                });
            }
        }
    }
}
