-- +goose Up
-- +goose StatementBegin
CREATE TABLE inventories (
  inventory_code VARCHAR(16) PRIMARY KEY,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_inventory_code ON inventories (inventory_code);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS inventories;
-- +goose StatementEnd
