export interface Message {
  id: number;
  content: string;
  hug_count: number;
  created_at: string;
}

export interface MessageListResponse {
  messages: Message[];
  next_cursor?: number;
}

export interface CreateMessageRequest {
  content: string;
}
