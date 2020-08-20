
from module_a import module_a

class ModuleB:

    def __init__(self):
        print("Init module b")
        self._name = "module_b"

    def do_something(self):
        mod_a = module_a.ModuleA()
        print(mod_a.do_something())
        return "{} is doing something".format(self._name)
