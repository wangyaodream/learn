import unittest

from common.cart import ShoppingCart
from common.product import Product


class ShoppingCartTestCase(unittest.TestCase):
    def test_add_and_remove_product(self):
        cart = ShoppingCart()
        product = Product('shoes', 'S', 'blue')

        cart.add_product(product, 10)
        cart.remove_product(product, 10)

        self.assertEqual({}, cart.products)
