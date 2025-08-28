CREATE TYPE transaction_type AS ENUM ('CREDIT', 'DEBIT', 'TOP_UP');
CREATE TYPE currency_type AS ENUM ('SAR', 'USD');

CREATE TABLE accounts (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_name VARCHAR(255) NOT NULL,
    balance BIGINT NOT NULL DEFAULT 0,
    currency currency_type NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);

CREATE TABLE transactions (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_uuid UUID NOT NULL REFERENCES accounts(uuid),
    idempotency_id VARCHAR(36) NOT NULL UNIQUE,
    trx_type transaction_type NOT NULL,
    amount BIGINT NOT NULL,
    after_balance BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);

CREATE INDEX idx_transactions_create_at ON transactions (created_at);
CREATE INDEX idx_transactions_account_uuid ON transactions (account_uuid);

ALTER TABLE accounts ADD CONSTRAINT check_balance_non_negative CHECK (balance >= 0);
ALTER TABLE transactions ADD CONSTRAINT check_after_balance_non_negative CHECK (after_balance >= 0);

