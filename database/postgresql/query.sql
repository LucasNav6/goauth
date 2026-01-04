-- name: GetUser :one
SELECT * FROM "user"
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM "user"
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY createdAt;

-- name: CreateUser :one
INSERT INTO "user" (
    id, name, email, emailVerified, image, createdAt, updatedAt
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE "user"
SET name = $2,
        email = $3,
        emailVerified = $4,
        image = $5,
        updatedAt = $6
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;

-- name: GetSession :one
SELECT * FROM session
WHERE id = $1
LIMIT 1;

-- name: GetSessionByToken :one
SELECT * FROM session
WHERE token = $1
LIMIT 1;

-- name: ListSessions :many
SELECT * FROM session
ORDER BY createdAt;

-- name: ListSessionsByUserId :many
SELECT * FROM session
WHERE userId = $1
ORDER BY createdAt;

-- name: CreateSession :one
INSERT INTO session (
    id, userId, token, expiresAt, ipAddress, userAgent, createdAt, updatedAt
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateSession :exec
UPDATE session
SET userId = $2,
        token = $3,
        expiresAt = $4,
        ipAddress = $5,
        userAgent = $6,
        updatedAt = $7
WHERE id = $1;

-- name: DeleteSession :exec
DELETE FROM session
WHERE id = $1;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1
LIMIT 1;

-- name: GetAccountByProviderAndUserId :one
SELECT * FROM account
WHERE providerId = $1
    AND userId = $2
LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY createdAt;

-- name: ListAccountsByUserId :many
SELECT * FROM account
WHERE userId = $1
ORDER BY createdAt;

-- name: CreateAccount :one
INSERT INTO account (
    id, userId, accountId, providerId, accessToken, refreshToken,
    accessTokenExpiresAt, refreshTokenExpiresAt, scope, idToken, password, createdAt, updatedAt
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING *;

-- name: UpdateAccount :exec
UPDATE account
SET userId = $2,
        accountId = $3,
        providerId = $4,
        accessToken = $5,
        refreshToken = $6,
        accessTokenExpiresAt = $7,
        refreshTokenExpiresAt = $8,
        scope = $9,
        idToken = $10,
        password = $11,
        updatedAt = $12
WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;

-- name: GetVerification :one
SELECT * FROM verification
WHERE id = $1
LIMIT 1;

-- name: GetVerificationByIdentifier :many
SELECT * FROM verification
WHERE identifier = $1
ORDER BY expiresAt;

-- name: ListVerifications :many
SELECT * FROM verification
ORDER BY createdAt;

-- name: CreateVerification :one
INSERT INTO verification (
    id, identifier, value, expiresAt, createdAt, updatedAt
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateVerification :exec
UPDATE verification
SET identifier = $2,
        value = $3,
        expiresAt = $4,
        updatedAt = $5
WHERE id = $1;

-- name: DeleteVerification :exec
DELETE FROM verification
WHERE id = $1;