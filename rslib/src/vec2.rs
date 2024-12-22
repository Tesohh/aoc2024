use std::fmt::Display;

#[derive(Debug, Hash, Copy, Clone, PartialEq, Eq, PartialOrd, Ord)]
pub struct Vec2 {
    pub x: i64,
    pub y: i64,
}

impl Display for Vec2 {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Vec2 {{x:{}, y:{}}}", self.x, self.y)
    }
}

impl Vec2 {
    pub const ZERO: Vec2 = Vec2 { x: 0, y: 0 };

    pub const UP: Vec2 = Vec2 { x: 0, y: -1 };
    pub const DOWN: Vec2 = Vec2 { x: 0, y: 1 };
    pub const LEFT: Vec2 = Vec2 { x: -1, y: 0 };
    pub const RIGHT: Vec2 = Vec2 { x: 1, y: 0 };

    pub const DIRECTIONS: [Vec2; 4] = [Vec2::UP, Vec2::DOWN, Vec2::LEFT, Vec2::RIGHT];

    pub const fn new(x: i64, y: i64) -> Self {
        Self { x, y }
    }

    pub const fn splat(n: i64) -> Self {
        Self { x: n, y: n }
    }

    pub fn add(self, other: Self) -> Self {
        Self {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }
    pub fn sub(self, other: Self) -> Self {
        Self {
            x: self.x - other.x,
            y: self.y - other.y,
        }
    }
    pub fn mul(self, other: Self) -> Self {
        Self {
            x: self.x * other.x,
            y: self.y * other.y,
        }
    }
    pub fn div(self, other: Self) -> Self {
        Self {
            x: self.x / other.x,
            y: self.y / other.y,
        }
    }

    pub fn as_direction_str(self) -> String {
        match self {
            Self::UP => String::from("Up"),
            Self::DOWN => String::from("Down"),
            Self::LEFT => String::from("Left"),
            Self::RIGHT => String::from("Right"),
            Self::ZERO => String::from("Zero"),
            _ => String::from("unknown"),
        }
    }
    pub fn as_direction_char(self) -> char {
        match self {
            Self::UP => '^',
            Self::DOWN => 'v',
            Self::LEFT => '<',
            Self::RIGHT => '>',
            Self::ZERO => '@',
            _ => '?',
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_arithmetic() {
        let p1 = Vec2::new(2, 3);
        let p2 = p1.add(Vec2::splat(4));
        let p3 = p1.sub(Vec2::splat(2));

        assert_eq!(p1, Vec2::new(2, 3));
        assert_eq!(p2, Vec2::new(6, 7));
        assert_eq!(p3, Vec2::new(0, 1));
    }
}
