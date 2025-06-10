CREATE TABLE IF NOT EXISTS places (
                                      id BIGSERIAL PRIMARY KEY,
                                      name VARCHAR(255) NOT NULL,
                                      type VARCHAR(100),
                                      address VARCHAR(255),
                                      city VARCHAR(100),
                                      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);