package store

var standardTax = newTaxRate(0.25, 20)

type Product struct {
    Name, Category string
    price float64
}

func NewProduct(name, category string, price float64) *Product {
    return &Product{ name, category, price }
}

func (p *Product) Price() float64 {
    // 由于是模块内调用，所以可以使用calcTax方法。calcTax是一个私有方法
    return standardTax.calcTax(p)
}

func (p *Product) SetPrice(newPrice float64) {
    p.price = newPrice
}
