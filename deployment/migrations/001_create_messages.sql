-- 创建 messages 表
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL CHECK (char_length(content) > 0 AND char_length(content) <= 500),
    hug_count INTEGER NOT NULL DEFAULT 0 CHECK (hug_count >= 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ip_address VARCHAR(45) NOT NULL
);

-- 添加表注释
COMMENT ON TABLE messages IS '树洞留言表';
COMMENT ON COLUMN messages.id IS '留言ID';
COMMENT ON COLUMN messages.content IS '留言内容（1-500字）';
COMMENT ON COLUMN messages.hug_count IS '抱抱次数';
COMMENT ON COLUMN messages.created_at IS '创建时间';
COMMENT ON COLUMN messages.ip_address IS 'IP地址（用于防刷）';
