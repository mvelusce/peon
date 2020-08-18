import unittest
from module_a import ModuleA


class TestModuleA(unittest.TestCase):

    def test_do_something(self):
        m = ModuleA()
        res = m.do_something()
        print(res)
        self.assertEqual(res, "module_a is doing something")

if __name__ == '__main__':
    unittest.main()
