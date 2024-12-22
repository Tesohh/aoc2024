use crate::Vec2;

// assumes all rows are of same length.
pub type Grid<T> = Vec<Vec<T>>;

pub trait GridExt<T> {
    fn width(&self) -> usize;
    fn height(&self) -> usize;

    fn is_out(&self, x: usize, y: usize) -> bool;
    fn is_out_point(&self, point: Vec2) -> bool;

    fn at(&self, x: usize, y: usize) -> &T;
    fn at_mut(&mut self, x: usize, y: usize) -> &mut T;

    fn at_point(&self, point: Vec2) -> &T; // panics if value doesn't exist
    fn at_point_mut(&mut self, point: Vec2) -> &mut T; // panics if value doesn't exist

    fn safe_at(&self, x: usize, y: usize) -> Option<&T>;
    fn safe_at_mut(&mut self, x: usize, y: usize) -> Option<&mut T>;
}

impl<T> GridExt<T> for Vec<Vec<T>> {
    fn width(&self) -> usize {
        self.len()
    }

    fn height(&self) -> usize {
        self[0].len()
    }

    fn is_out(&self, x: usize, y: usize) -> bool {
        let x_out = x > self.height() - 1;
        let y_out = y > self.width() - 1;

        x_out || y_out
    }

    fn is_out_point(&self, point: Vec2) -> bool {
        if point.x < 0 || point.y < 0 {
            return true;
        }
        self.is_out(point.x as usize, point.y as usize)
    }

    fn at(&self, x: usize, y: usize) -> &T {
        self.safe_at(x, y).unwrap()
    }

    fn at_mut(&mut self, x: usize, y: usize) -> &mut T {
        self.safe_at_mut(x, y).unwrap()
    }

    fn at_point(&self, point: Vec2) -> &T {
        self.safe_at(point.x as usize, point.y as usize).unwrap()
    }

    fn at_point_mut(&mut self, point: Vec2) -> &mut T {
        self.safe_at_mut(point.x as usize, point.y as usize)
            .unwrap()
    }

    fn safe_at(&self, x: usize, y: usize) -> Option<&T> {
        self.get(y as usize)?.get(x as usize)
    }

    fn safe_at_mut(&mut self, x: usize, y: usize) -> Option<&mut T> {
        self.get_mut(y as usize)?.get_mut(x as usize)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_safe_at() {
        let grid: Grid<usize> = vec![
            vec![00, 01, 02, 03, 04, 05],
            vec![10, 11, 12, 13, 14, 15],
            vec![20, 21, 22, 23, 24, 25],
        ];

        assert_eq!(grid.safe_at(100, 100), None);
        assert_eq!(grid.safe_at(2, 1), Some(&12));
    }

    #[test]
    fn test_at_point() {
        let grid: Grid<usize> = vec![
            vec![00, 01, 02, 03, 04, 05],
            vec![10, 11, 12, 13, 14, 15],
            vec![20, 21, 22, 23, 24, 25],
        ];

        assert_eq!(grid.at_point(Vec2::new(2, 1)), &12usize);
        assert!(std::panic::catch_unwind(|| grid.at_point(Vec2::new(100, 100))).is_err())
    }
}
