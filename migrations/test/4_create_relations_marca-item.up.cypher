MATCH
  (a:Item),
  (b:Marca)
  WHERE a.sku = 33 AND b.nome = 'Nike'
CREATE (a)-[r:DI]->(b);


MATCH
  (a:Item),
  (b:Marca)
  WHERE a.sku = 11 AND b.nome = 'Nike'
CREATE (a)-[r:DI]->(b);

MATCH
  (a:Item),
  (b:Marca)
  WHERE a.sku = 22 AND b.nome = 'Nike'
CREATE (a)-[r:DI]->(b);

MATCH
  (a:Item),
  (b:Marca)
  WHERE a.sku = 44 AND b.nome = 'Nike'
CREATE (a)-[r:DI]->(b);

MATCH
  (a:Item),
  (b:Marca)
  WHERE a.sku = 55 AND b.nome = 'Adidas'
CREATE (a)-[r:DI]->(b);
