use aoc::Vec2;

#[derive(Debug, Clone)]
pub struct Path {
    pub cells: Vec<Vec2>,
    pub cost: i32,

    pub dir: Vec2,
    pub retired: bool,
    pub finished: bool,
}
