CREATE TABLE match_cats (
    matchId VARCHAR(255) PRIMARY KEY,
    matchCatId VARCHAR(255),
    userCatId VARCHAR(255),
    status VARCHAR(10),
    issuedBy VARCHAR(255),
    message VARCHAR(255),
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)