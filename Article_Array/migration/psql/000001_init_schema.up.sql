CREATE TABLE IF NOT EXISTS "author"(
    "id" SERIAL PRIMARY KEY,
    "firstname" VARCHAR(255) NOT NULL,
    "lastname" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP DEFAULT(NOW()),
    "updated_at" TIMESTAMP DEFAULT(NOW()),
    "deleted_at" TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS "article"(
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "body" VARCHAR(255) NOT NULL,
    "author_id" INT,
    "created_at" TIMESTAMP DEFAULT(NOW()),
    "updated_at" TIMESTAMP DEFAULT(NOW()),
      "deleted_at" TIMESTAMP NULL,
    CONSTRAINT fk_author FOREIGN KEY(author_id) REFERENCES author(id)
);