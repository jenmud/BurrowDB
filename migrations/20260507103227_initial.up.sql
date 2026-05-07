CREATE TABLE IF NOT EXISTS nodes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    properties JSONB NOT NULL DEFAULT '{}'::JSONB
);


CREATE INDEX IF NOT EXISTS node_property_idx ON nodes USING GIN(properties);


CREATE TABLE IF NOT EXISTS node_labels (
    name VARCHAR(255) NOT NULL,
    node_id INTEGER REFERENCES nodes(id)
);


CREATE TABLE IF NOT EXISTS edges (
    id SERIAL PRIMARY KEY,
    source_id INTEGER REFERENCES nodes(id),
    target_id INTEGER REFERENCES nodes(id),
    properties JSONB NOT NULL DEFAULT '{}'::JSONB
);


CREATE INDEX IF NOT EXISTS edge_property_idx ON edges USING GIN(properties);
CREATE INDEX IF NOT EXISTS edge_source_idx ON edges(source_id);
CREATE INDEX IF NOT EXISTS edge_target_idx ON edges(target_id);
CREATE INDEX IF NOT EXISTS edge_outgoing_idx ON edges(source_id, target_id);
CREATE INDEX IF NOT EXISTS edge_incoming_idx ON edges(target_id, source_id);


CREATE TABLE IF NOT EXISTS edge_labels (
    name VARCHAR(255) NOT NULL,
    edge_id INTEGER REFERENCES edges(id)
);
