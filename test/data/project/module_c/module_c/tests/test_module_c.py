import unittest
from module_c import module_c


class TestModuleC(unittest.TestCase):

    def test_do_something(self):
        m = module_c.ModuleC()
        res = m.do_something()
        print(res)
        self.assertEqual(res, "module_c is doing something")

if __name__ == '__main__':
    unittest.main()
