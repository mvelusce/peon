
class ModuleA:

    def __init__(self):
        print("Init module a")
        self._name = "module_a"

    def do_something(self):
        return "{} is doing something".format(self._name)
