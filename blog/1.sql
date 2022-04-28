DROP DATABASE IF exists blog;

CREATE DATABASE blog charset=UTF8;

USE blog;

CREATE TABLE users(
    id INT ,-- 记录用户数目
    account VARCHAR(10) PRIMARY KEY,-- 用户账号
    password VARCHAR(15),
    login_time VARCHAR(30)
);

CREATE TABLE stars(
    account VARCHAR(10) PRIMARY KEY,-- 用户账号
    name TEXT,
    time VARCHAR(30),
    CONSTRAINT first 
    FOREIGN KEY(account) REFERENCES users(account)
)