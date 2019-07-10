import gaussian_elimination


class Vector:
    """二维空间的向量类"""

    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __add__(self, other):
        """向量的加法，就是向量的各个坐标相加"""
        return Vector(self.x + other.x, self.y + other.y)

    def __mul__(self, other):
        """向量与标量的乘法，就是向量的各个坐标分别乘以标量"""
        if isinstance(other, int):
            return Vector(self.x * other, self.y + other)
        else:
            raise Exception("暂时不支持其他类型和向量相乘")

    def __str__(self):
        return "({},{})".format(self.x, self.y)


def is_inear_combination(a1, a2, b):
    """判断b是否是向量a1,a2的线性组合"""
    ##TODO 首先判断abc是否列元素相等，假设已经判断好了
    a = [[a1.x, a2.x, b.x],
         [a1.y, a2.y, b.y]
         ]
    gaussian_elimination.gaussian_elimination(a)
    gaussian_elimination.print_matrix(a)
    pass
