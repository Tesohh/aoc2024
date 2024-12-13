def cramer(d_ax: int, d_bx: int, d_ay: int, d_by: int, expect_x: int, expect_y: int):
    """
        solves this system:
        { delta_a_x * n + delta_b_x * m = expect_x
        { delta_a_y * n + delta_b_y * m = expect_y
    """
    det = (d_ax * d_by) - (d_bx * d_ay)

    n = ((expect_x * d_by) - (expect_y * d_bx)) / det
    m = ((d_ax * expect_y) - (d_ay * expect_x)) / det 

    return (n, m)

print(cramer(3, 2, 2, 1, 43, 28))
