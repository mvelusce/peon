
from module_a import module_a
from module_b import module_b

class ModuleC:

    def __init__(self):
        print("Init module c")
        self._name = "module_c"

    def do_something(self):
        mod_a = module_a.ModuleA()
        print(mod_a.do_something())
        mod_b = module_b.ModuleB()
        print(mod_b.do_something())
        return "{} is doing something".format(self._name)
