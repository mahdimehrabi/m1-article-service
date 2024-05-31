CREATE TABLE IF NOT EXISTS articles (
                                    id BIGSERIAL PRIMARY KEY,
                                    title varchar(50) NOT NULL,
                                    slug varchar(100) NOT NULL,
                                    tags varchar(30)[] NOT NULL,
                                    created_at BIGINT NOT NULL
);
