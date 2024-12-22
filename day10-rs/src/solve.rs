use std::collections::HashSet;

use aoc::{GridExt, Vec2};

use crate::hiker::Hiker;

pub fn solve(grid: aoc::Grid<i64>) -> (usize, usize) {
    let mut trailheads: Vec<Hiker> = vec![];

    for y in 0..grid.height() {
        for x in 0..grid.width() {
            if *grid.at(x, y) == 0 {
                trailheads.push(Hiker {
                    pos: aoc::Vec2::new(x as i64, y as i64),
                    dir: aoc::Vec2::ZERO,
                    retired: false,
                });
            }
        }
    }

    let mut score = 0;
    let mut rating = 0;

    for trailhead in trailheads {
        let mut peaks: HashSet<Vec2> = HashSet::new();
        let mut hikers: Vec<Hiker> = vec![trailhead];

        while hikers.len() > 0 {
            let mut new_hikers: Vec<Hiker> = vec![];
            hikers = hikers.into_iter().filter(|h| !h.retired).collect();

            for hiker in &mut hikers {
                let current = *grid.at_point(hiker.pos);

                // figure out possible branches
                let mut branches: Vec<Vec2> = vec![];
                for dir in Vec2::DIRECTIONS {
                    let next_point = hiker.pos.add(dir);

                    if grid.is_out_point(next_point) {
                        continue;
                    }

                    if *grid.at_point(next_point) == current + 1 {
                        branches.push(dir);
                    }
                }

                // manage branching
                if branches.len() == 0 {
                    hiker.retired = true
                } else if branches.len() == 1 {
                    hiker.dir = branches[0]
                } else {
                    hiker.retired = true;
                    for dir in branches {
                        new_hikers.push(Hiker {
                            pos: hiker.pos,
                            dir,
                            retired: false,
                        });
                    }
                }
            }

            hikers.extend(new_hikers);

            // move forth
            for hiker in &mut hikers {
                if hiker.retired == true {
                    continue;
                }

                let next_point = hiker.pos.add(hiker.dir);
                if grid.is_out_point(next_point) {
                    hiker.retired = true;
                    continue;
                }

                hiker.pos = next_point;

                if *grid.at_point(hiker.pos) == 9 {
                    hiker.retired = true;
                    rating += 1;
                    peaks.insert(hiker.pos);
                }
            }
        }

        score += peaks.len();
    }

    (score, rating)
}
