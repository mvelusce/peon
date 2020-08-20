import unittest
from module_b import module_b


class TestModuleB(unittest.TestCase):

    def test_do_something(self):
        m = module_b.ModuleB()
        res = m.do_something()
        print(res)
        self.assertEqual(res, "module_b is doing something")

if __name__ == '__main__':
    unittest.main()
