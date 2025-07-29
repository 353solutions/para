class Item:
    def __init__(self, x, y):
        self.x, self.y = x, y

    def move(self, x, y):
        print('type:', type(self))
        self.x, self.y = x, y


class Player(Item):
    def __init__(self, name, x, y):
        self.name = name
        super().__init__(x, y)


p1 = Player('Parzival', 100, 300)
p1.move(400, 500)


# Not possible in Go: extending
class Driver:
    def execute(self, query, params):
        msg = self.prepare(query, params)
        self.conn.call(msg)
        out = self.conn.read()
        return self.parse(out)

    def prepare(self, query, params):
        raise NotImplementedError

    def parse(self, out):
        raise NotImplementedError


class PGDriver(Driver):
    def prepare(self, query, params):
        return ''
