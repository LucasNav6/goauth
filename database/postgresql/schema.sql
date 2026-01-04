CREATE TABLE IF NOT EXISTS "user" (
    id TEXT PRIMARY KEY,
    name TEXT,
    email TEXT UNIQUE,
    emailVerified BOOLEAN NOT NULL DEFAULT FALSE,
    image TEXT,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS session (
    id TEXT PRIMARY KEY,
    userId TEXT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    expiresAt TIMESTAMPTZ NOT NULL,
    ipAddress TEXT,
    userAgent TEXT,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS account (
    id TEXT PRIMARY KEY,
    userId TEXT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    accountId TEXT NOT NULL,
    providerId TEXT NOT NULL,
    accessToken TEXT,
    refreshToken TEXT,
    accessTokenExpiresAt TIMESTAMPTZ,
    refreshTokenExpiresAt TIMESTAMPTZ,
    scope TEXT,
    idToken TEXT,
    password TEXT,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT account_provider_accountid_unique UNIQUE (providerId, accountId)
);

CREATE TABLE IF NOT EXISTS verification (
    id TEXT PRIMARY KEY,
    identifier TEXT NOT NULL,
    value TEXT NOT NULL,
    expiresAt TIMESTAMPTZ NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_verification_identifier ON verification(identifier);
CREATE INDEX IF NOT EXISTS idx_session_userid ON session(userId);
CREATE INDEX IF NOT EXISTS idx_account_userid ON account(userId);