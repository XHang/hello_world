class user:
    def __init__(self):
        self.name = ""

    @property
    def name(self):
        print("即将获取", "  name值")
        return self._name

    @name.setter
    def name(self, name):
        print("即将设置", "name值")
        self._name = name


u = user()
u.name = "王兰花"
print(u.name)
