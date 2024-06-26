CREATE TABLE cats (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    race VARCHAR(50) NOT NULL,
    sex VARCHAR(6) NOT NULL,
    ageInMonths INTEGER NOT NULL,
    description VARCHAR(255) NOT NULL,
    hasMatched BOOLEAN NOT NULL DEFAULT FALSE,
    userId VARCHAR(255) NOT NULL,
    imageUrls VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);