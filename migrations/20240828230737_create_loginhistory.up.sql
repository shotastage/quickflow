-- Create loginhistory table
CREATE TABLE loginhistory (
    id TEXT PRIMARY KEY AUTOINCREMENT NOT NULL,
    userid TEXT NOT NULL,
    ipaddress VARCHAR(255) NOT NULL,
    useragent VARCHAR(255) NOT NULL,
    logintime TEXT NOT NULL
);
