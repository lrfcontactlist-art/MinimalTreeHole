-- 创建索引以优化查询性能

-- 按创建时间倒序查询的索引（主要查询场景）
CREATE INDEX IF NOT EXISTS idx_messages_created_at_desc ON messages (created_at DESC);

-- IP地址索引（用于防刷检查）
CREATE INDEX IF NOT EXISTS idx_messages_ip_address ON messages (ip_address);

-- 复合索引：IP + 创建时间（用于"同一IP在X分钟内发布次数"查询）
CREATE INDEX IF NOT EXISTS idx_messages_ip_created ON messages (ip_address, created_at DESC);
