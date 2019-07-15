import gaussian_elimination


class Vector:
    """二维空间的向量类"""

    def __init__(self, *x):
        self.elements = x

    def __add__(self, other):
        """向量的加法，就是向量的各个坐标相加"""
        if len(self.elements) != len(other.elements):
            raise Exception("向量元素大小不一致，不能相加")
        result = [self.elements[x] + other.elements[x] for x in range(len(self.elements))]

        return Vector(result)

    def __mul__(self, other):
        """向量与标量的乘法，就是向量的各个坐标分别乘以标量"""
        if isinstance(other, int):
            return Vector([x * other for x in self.elements])
        else:
            raise Exception("暂时不支持其他类型和向量相乘")

    def __str__(self):
        str = "("
        for x in self.elements:
            str += "{},".format(x)
        str = str[:len(str) - 1]
        str += ")"
        return str

    def __len__(self):
        return len(self.elements)

    def __getitem__(self, item):
        return self.elements[item]


def is_inear_combination(a1, a2, b):
    if len(a1) != len(a2) or len(a1) != len(b):
        return False
    matrix = []
    for index in range(len(a1)):
        matrix.append([a1[index], a2[index], b[index]])
    result = gaussian_elimination.gaussian_elimination(matrix)
    print("原始矩阵")
    gaussian_elimination.print_matrix(matrix)
    print("化简矩阵")
    gaussian_elimination.print_matrix(result)
