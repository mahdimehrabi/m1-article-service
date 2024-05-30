CREATE TABLE IF NOT EXISTS articles (
                                    id BIGSERIAL PRIMARY KEY,
                                    title string NOT NULL,
                                    slug string NOT NULL,
                                    tags string NOT NULL,
                                    created_at BIGINT NOT NULL,
);
