import unittest
from module_d import module_d


class TestModuleD(unittest.TestCase):

    def test_do_something(self):
        m = module_d.ModuleD()
        res = m.do_something()
        print(res)
        self.assertEqual(res, "module_d is doing something")

if __name__ == '__main__':
    unittest.main()
