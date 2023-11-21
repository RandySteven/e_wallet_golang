DROP DATABASE e_wallet_db;
CREATE DATABASE e_wallet_db;

DROP TABLE IF EXISTS user_profiles;
DROP TABLE IF EXISTS transfer_transactions;
DROP TABLE IF EXISTS topup_transactions;
DROP TABLE IF EXISTS games;
DROP TABLE IF EXISTS boxes;
DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS forgot_passwords;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    birthday DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_profiles(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    user_id BIGINT NOT NULL,
    password VARCHAR NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS forgot_passwords(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    new_password VARCHAR NOT NULL,
    reset_password_token VARCHAR NOT NULL,
    expiry_time TIME NOT NULL,
    is_valid BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS wallets (
    id BIGSERIAL PRIMARY KEY,
    balance INT NOT NULL,
    wallet_number VARCHAR NOT NULL UNIQUE,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS transfer_transactions(
    id BIGSERIAL PRIMARY KEY,
    sender_wallet_id BIGINT NOT NULL,
    receiver_wallet_id BIGINT NOT NULL,
    amount INT NOT NULL,
    description VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY(sender_wallet_id) REFERENCES wallets(id),
    FOREIGN KEY(receiver_wallet_id) REFERENCES wallets(id)
);

CREATE TABLE IF NOT EXISTS topup_transactions (
    id BIGSERIAL PRIMARY KEY,
    wallet_id BIGINT NOT NULL,
    amount INT NOT NULL,
    source_of_fund VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY(wallet_id) REFERENCES wallets(id)
);

CREATE TABLE IF NOT EXISTS boxes (
    id BIGSERIAL PRIMARY KEY,
    amount INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS games(
    id BIGSERIAL PRIMARY KEY,
    wallet_id BIGINT NOT NULL,
    box_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY(wallet_id) REFERENCES wallets(id),
    FOREIGN KEY(box_id) REFERENCES boxes(id)
);

-- 5 users
INSERT INTO users (name, birthday)
VALUES
    ('Randy Steven', '11-04-2001'),
    ('Don Joe', '12-01-1999'),
    ('Silver Bullet', '21-03-2000'),
    ('Big Mom', '22-04-1992'),
    ('Meme Mimi Momo', '20-05-2002');

INSERT INTO user_profiles (user_id, name, email, password)
VALUES
    (1, 'randy_steven', 'randy.steven@shopee.com', 'test_1234'),
    (2, 'don_joe', 'don.joe@shopee.com', 'test_2345'),
    (3, 'peluru_perak', 'silver.bullet@shopee.com', 'test_3456'),
    (4, 'big_mama', 'big.mom@shopee.com', 'test_4567'),
    (5, 'momomimi', 'mimo.mimo@shopee.com', 'test_4567');

INSERT INTO wallets (user_id, balance, wallet_number)
VALUES
    (1, 5000000, 'WR000001'),
    (2, 3000000, 'WR000002'),
    (3, 4500000, 'WR000003'),
    (4, 5000000, 'WR000004'),
    (5, 5200000, 'WR000005');


INSERT INTO transfer_transactions (sender_wallet_id, receiver_wallet_id, amount, description, created_at)
VALUES
    (1, 2, 20000, 'Cicilan utang', '2022-01-10'),
    (1, 3, 100000, 'Pinjam dulu seratus', '2022-01-11'),
    (3, 1, 110000, 'Balikin dulu seratus', '2022-02-03'),
    (2, 4, 52000, 'Beli kartu', '2022-03-23'),
    (4, 5, 25000, 'Titipkan dari pak cik', '2022-05-04'),
    (3, 4, 14000, 'Beli es teh manis', '2022-06-09'),
    (5, 1, 133000, 'Beli sembako di pasar', '2022-06-24'),
    (1, 3, 51000, '50rb bang', '2022-07-22'),
    (2, 1, 145000, 'Tidak tau lah', '2022-07-27'),
    (1, 2, 10400, 'Tidak tau lah', '2022-08-19'),
    (1, 2, 20560, 'Cicilan utang', '2022-10-11'),
    (1, 3, 100300, 'Pinjam dulu seratus', '2022-10-12'),
    (3, 1, 103000, 'Balikin dulu seratus', '2022-11-16'),
    (2, 5, 10220, 'Beli kartu', '2022-11-18'),
    (4, 5, 25060, 'Titipkan dari pak cik', '2023-03-24'),
    (3, 4, 10066, 'Beli es teh manis', '2023-03-27'),
    (5, 1, 122200, 'Beli sembako di pasar', '2023-04-04'),
    (1, 3, 54320, '50rb bang', '2023-04-12'),
    (2, 1, 121400, 'Tidak tau lah', '2023-05-19'),
    (1, 2, 34500, 'Tidak tau lah', '2023-07-11'),
    (3, 5, 16900, 'Beli es teh manis', '2023-07-12'),
    (5, 1, 166600, 'Beli sembako di pasar', '2023-07-28'),
    (1, 3, 53210, '50rb bang', '2023-08-09'),
    (2, 1, 100004, 'Tidak tau lah', '2023-08-11'),
    (5, 2, 30005, 'Tidak tau lah', '2023-08-24');

INSERT INTO topup_transactions (wallet_id, amount, source_of_fund, created_at)
VALUES
    (1, 20001, 'DEBIT', '2022-01-10'),
    (1, 100000, 'CREDIT', '2022-01-11'),
    (3, 100000, 'DEBIT', '2022-02-03'),
    (2, 50000, 'REWARD', '2022-03-23'),
    (4, 25005, 'REWARD', '2022-05-04'),
    (3, 10012, 'CREDIT', '2022-06-09'),
    (5, 100014, 'DEBIT', '2022-06-24'),
    (1, 50009, 'DEBIT', '2022-07-22'),
    (2, 100123, 'CREDIT', '2022-07-27'),
    (1, 30444, 'REWARD', '2022-08-19'),
    (5, 20999, 'DEBIT', '2022-10-11'),
    (1, 100009, 'CREDIT', '2022-10-12'),
    (5, 100021, 'DEBIT', '2022-11-16'),
    (5, 50032, 'REWARD', '2022-11-18'),
    (4, 25456, 'REWARD', '2023-03-24'),
    (2, 10678, 'CREDIT', '2023-03-27'),
    (5, 131313, 'DEBIT', '2023-04-04'),
    (1, 51000, 'DEBIT', '2023-04-12'),
    (2, 199200, 'CREDIT', '2023-05-19'),
    (1, 31230, 'REWARD', '2023-07-11'),
    (2, 18880, 'CREDIT', '2023-07-12'),
    (5, 100005, 'DEBIT', '2023-07-28'),
    (1, 50070, 'DEBIT', '2023-08-09'),
    (2, 100002, 'CREDIT', '2023-08-11'),
    (1, 30001, 'REWARD', '2023-08-24');

INSERT INTO boxes (amount)
VALUES
    (10000),
    (11000),
    (12000),
    (13000),
    (14000),
    (15000),
    (16000),
    (17000),
    (18000),
    (19000),
    (20000);

SELECT u.name, up.email, w.number, w.balance FROM users u JOIN user_profiles up
ON u.id = up.user_id JOIN wallets w
ON u.id = w.user_id;

UPDATE wallets 
SET balance = 100000;