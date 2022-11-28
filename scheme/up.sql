CREATE ROLE dns_vladivostok WITH PASSWORD '12345' LOGIN;

CREATE DATABASE dns_assigment WITH OWNER dns_Vladivostok;

CREATE TABLE map(
    key INT,
    value INT,
    UNIQUE(key)
);