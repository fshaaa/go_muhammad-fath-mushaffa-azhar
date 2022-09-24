-- QUERY in Redis, Neo4J, Cassandra
-- SELECT * FROM users;

-- Redis
HGETALL USERS

-- Neo4J
MATCH users
RETURN *

-- Cassandra
SELECT * FROM users
