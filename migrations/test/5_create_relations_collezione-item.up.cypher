MATCH
  (a:Item),
  (b:Collezione)
  WHERE a.sku = 33 AND b.nome = 'Estate'
CREATE (a)-[:Appartiene]->(b)
WITH true as pass

MATCH
  (a:Item),
  (b:Collezione)
  WHERE a.sku = 11 AND b.nome = 'Estate'
CREATE (a)-[:Appartiene]->(b)
WITH true as pass

MATCH
  (a:Item),
  (b:Collezione)
  WHERE a.sku = 22 AND b.nome = 'Inverno'
CREATE (a)-[:Appartiene]->(b)
WITH true as pass

MATCH
  (a:Item),
  (b:Collezione)
  WHERE a.sku = 44 AND b.nome = 'Inverno'
CREATE (a)-[:Appartiene]->(b)
WITH true as pass

MATCH
  (a:Item),
  (b:Collezione)
  WHERE a.sku = 55 AND b.nome = 'Estate'
CREATE (a)-[:Appartiene]->(b)
