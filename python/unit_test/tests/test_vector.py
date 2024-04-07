import sys
import unittest

from vector import Vector


class TestVector(unittest.TestCase):
    # 在方法前打印开始和结束
    # def setUp(self) -> None:
    #     print("start")
    #
    # def tearDown(self) -> None:
    #     print("end")

    # 在类中打印
    # @classmethod
    # def setUpClass(cls) -> None:
    #     print("start")
    #
    # @classmethod
    # def tearDownClass(cls) -> None:
    #     print("end")

    @unittest.skipIf(sys.version_info < (3, 7) ,
                     "Do not support this version")
    def test_init(self):
        v = Vector(1, 2)
        self.assertEqual(v.x, 1)
        self.assertEqual(v.y, 2)

        with self.assertRaises(ValueError):
            v = Vector("1", "2")

    def test_add(self):
        v1 = Vector(1, 2)
        v2 = Vector(2, 3)
        v3 = v1.add(v2)
        self.assertEqual(v3.x , 3)
