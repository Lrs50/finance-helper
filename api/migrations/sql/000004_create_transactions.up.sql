CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    value NUMERIC NOT NULL CHECK (value > 0),
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE RESTRICT,
    comment VARCHAR(255),
    date TIMESTAMP NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_transactions_user_id ON transactions(user_id);
CREATE INDEX IF NOT EXISTS idx_transactions_date ON transactions(date);
CREATE INDEX IF NOT EXISTS idx_transactions_category_id ON transactions(category_id);
CREATE INDEX IF NOT EXISTS idx_transactions_tag_id ON transactions(tag_id);
