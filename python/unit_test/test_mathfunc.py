import unittest
from mathfunc import *


class TestMathFunc(unittest.TestCase):
    """Test mathfunc.py"""


    def test_add(self):
        """ Test method add(a, b)"""
        self.assertEqual(3, add(1, 2))
        self.assertNotEqual(3, add(2, 2))

    def test_minus(self):
        self.assertEqual(1, minus(3, 2))
        self.assertNotEqual(3, minus(2, 2))

    def test_multi(self):
        self.assertEqual(9, multi(3, 3))

    def test_divide(self):
        self.assertEqual(10, divide(50, 5))
        self.assertEqual(2.5, divide(5, 2))

    @unittest.skip("I don't want to run this case.")
    def test_skip(self):
        self.assertEqual(10, divide(50, 5))
        self.assertEqual(2.5, divide(5, 2))

if __name__ == "__main__":
    unittest.main(verbosity=2)
