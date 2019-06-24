class point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    # 重载字符串表现形式方法，print(instance)可见效果
    def __str__(self):
        return "self x is {} y is {}".format(self.x, self.y)

    # 重载+运算符，将两个对象+起来，将得到一个新的对象
    def __add__(self, other):
        return point(self.x + other.x, self.y + other.x)


p1 = point(1, 2)
p2 = point(3, 4)
print(p1)
print(p1 + p2)
