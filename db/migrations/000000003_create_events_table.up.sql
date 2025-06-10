CREATE TABLE IF NOT EXISTS events (
                                      id BIGSERIAL PRIMARY KEY,
                                      name VARCHAR(100) NOT NULL,
                                      artist VARCHAR(100),
                                      place_id BIGINT NOT NULL,
                                      date_time TIMESTAMPTZ NOT NULL,
                                      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                      CONSTRAINT fk_place
                                          FOREIGN KEY(place_id)
                                              REFERENCES places(id)
);