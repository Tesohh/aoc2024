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

    while paths.iter().filter(|p| !p.finished).count() > 0 {
        let mut new_paths: Vec<Path> = vec![];
        paths.retain(|p| !p.retired);

        for path in paths.iter_mut().filter(|p| !p.finished) {
            let try_directions: Vec<Vec2> = vec![
                path.dir,
                path.dir.turn_90deg_clockwise(),
                path.dir.turn_90deg_counter_clockwise(),
            ];

            println!("try_directions: {:?}", try_directions);
            let mut valid_directions: Vec<Vec2> = vec![];
            for dir in try_directions {
                let next = path.cells.last().unwrap().add(dir);
                if !grid.is_out_point(next) && *grid.at_point(next) != '#' {
                    valid_directions.push(dir);
                }
            }
            println!("valid_directions: {:?}", valid_directions);

            if valid_directions.is_empty() {
                path.retired = true
            } else {
                if !valid_directions.contains(&path.dir) {
                    path.retired = true
                }

                for dir in valid_directions {
                    let next = path.cells.last().unwrap().add(dir);

                    if dir == path.dir {
                        let mut new_path = path.clone();
                        new_path.cost += 1;
                        new_path.cells.push(next);
                        new_paths.push(new_path);
                    } else {
                        let mut new_path = path.clone();
                        new_path.cost += 1000;
                        new_path.dir = dir;
                        new_path.cells.push(next);
                        new_paths.push(new_path);
                    }

                    path.retired = true;

                    if *grid.at_point(next) == 'E' {
                        path.finished = true
                    }
                }
            }
        }
        paths.extend(new_paths);
        debug::debug_grid(grid.clone(), paths.clone());
    }

    dbg!(paths);
}
