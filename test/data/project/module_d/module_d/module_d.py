
from module_c import module_c
from module_b import module_b

class ModuleD:

    def __init__(self):
        print("Init module d")
        self._name = "module_d"

    def do_something(self):
        mod_c = module_c.ModuleC()
        print(mod_c.do_something())
        mod_b = module_b.ModuleB()
        print(mod_b.do_something())
        return "{} is doing something".format(self._name)
