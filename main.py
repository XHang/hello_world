import vector
from gaussian_elimination import gaussian_elimination
from vector import Vector

# a = Vector(1, -2, 3)
# b = Vector(5, -13, -3)
# c = Vector(-3, 8, 1)
# print("向量a{}和向量b{}相加的结果是{}".format(a, b, a + b))
#
# num = 3
# print("向量a{}和num{}相乘的结果是{}".format(a, num, a * 3))
# vector.is_inear_combination(a, b, c)

# 高斯化元法的使用
# 这个是引起BUG的数据源
# a = [
#     [-2, 8, 1, 0],
#     [3, 5, -6, 0],
# ]
# gaussian_elimination(a)
a = [
    [27.6, 30.2, 162],
    [3100, 6400, 23610],
    [250, 360, 1623]
]
gaussian_elimination(a)
