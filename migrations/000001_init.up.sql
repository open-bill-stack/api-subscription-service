BEGIN;

CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- для gen_random_uuid()

CREATE TABLE subscriptions
(
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id   UUID NOT NULL, -- ID користувача, який оформив підписку
    tariff_id UUID NOT NULL  -- ID тарифу
);


COMMIT;