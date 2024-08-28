-- Create shortenlink table
CREATE TABLE shortenlink (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    code VARCHAR(255),
    originalurl VARCHAR(255),
    isbanned BOOLEAN,
    securityreason VARCHAR(255),
    createdat TEXT,
    updatedat TEXT
);
