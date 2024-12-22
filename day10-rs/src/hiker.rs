#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub struct Hiker {
    pub pos: aoc::Vec2,
    pub dir: aoc::Vec2,
    pub retired: bool,
}
