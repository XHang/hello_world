import numpy
from fractions import Fraction


def swap(row, max_row, matrix):
    """
    交换矩阵中两个行的位置
    """
    if row is max_row:
        return
    temp = matrix[row]
    matrix[row] = matrix[max_row]
    matrix[max_row] = temp
    pass


# 说一下哈，高斯化元法致力于将主元列化为1，其他行化为0，所以是循环列的说
def gaussian_elimination(matrix):
    print("原始矩阵是")
    print_matrix(matrix)
    rows, cols = get_col_row(matrix)
    # 遍历列，从左到右
    row = 0
    for col in range(0, cols - 1):
        print("开始遍历第{}列,第{}行,矩阵是".format(col, row))
        print_matrix(matrix)
        max_row = get_maxrow_in_col(matrix, col, row)
        if matrix[max_row][col] == 0:
            print("矩阵的第{}列没有主元元素，跳过".format(col))
            continue
        print("第{}列最大的元素，在第{}行".format(col, max_row))
        print("交换{}行和{}行的矩阵后，矩阵是".format(row, max_row))
        swap(row, max_row, matrix)
        print_matrix(matrix)
        fraction = Fraction(1, matrix[row][col])
        row_multi(matrix, fraction, row)
        print("第{}行同乘以一个数{}".format(row, fraction))
        print_matrix(matrix)
        # 遍历剩下的行
        for i in range(row + 1, rows):
            if matrix[i][col] == 0:
                print("矩阵主元列{}第{}行元素已经是0，不用再化简".format(col, i))
                continue
            fraction = -Fraction(matrix[i][col], matrix[row][col])
            print("将第{}行的{}倍加到第{}行，并替代{}行，结果是".format(row, fraction, i, i))
            row_add_list(matrix, i, get_row_multiple(matrix, row, fraction))
            print_matrix(matrix)
        row = row + 1
    print("最后结果是")
    print_matrix(matrix)
    pass


def row_multi(matrix, num, row):
    for i in range(len(matrix[row])):
        result = matrix[row][i] * num
        if result.denominator == 1:
            matrix[row][i] = result.numerator
        else:
            matrix[row][i] = result


def row_add_list(matrix, row1, row_multi):
    """将list加到row1上面，并替代掉row1"""
    for col in range(len(matrix[0])):
        result = matrix[row1][col] + row_multi[col]
        if result.denominator == 1:
            matrix[row1][col] = result.numerator
        else:
            matrix[row1][col] = result


def get_row_multiple(matrix, row, multiple):
    elements = []
    for e in matrix[row]:
        result = e * multiple
        if result.denominator == 1:
            elements.append(result.numerator)
        else:
            elements.append(result)
    return elements


def get_col_row(matrix):
    rows = len(matrix)
    if rows is 0:
        return 0, 0
    return rows, len(matrix[0])


def print_matrix(matrix):
    for row in matrix:
        for col in row:
            if isinstance(col, Fraction):
                print(col.numerator, "/", col.denominator, end="           ", sep="")
            else:
                print(col, end="           ")
        print("")


# 找到指定列中，最大元素(绝对值)所在的行数
""" 
          1,2,3
 result  >4,5,6
          ^
         col
    return 1
"""


def get_maxrow_in_col(matrix, col, row_index):
    """从指定的列中拿到最大元素的行坐标（不包含比row_index更小的行）"""
    max_value = 0
    max_row = 0
    for row in range(row_index, len(matrix)):
        if abs(max_value) < abs(matrix[row][col]):
            max_value = matrix[row][col]
            max_row = row
    return max_row


a = [
    [1, 5, 7],
    [1, -2, -2]

]
gaussian_elimination(a)
