import { Message, MessageListResponse, CreateMessageRequest } from '../types';

const API_BASE_URL = '/api';

class APIClient {
  private baseURL: string;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  async createMessage(data: CreateMessageRequest): Promise<Message> {
    const response = await fetch(`${this.baseURL}/messages`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      let errorMessage = 'Failed to create message';
      try {
        const error = await response.json();
        errorMessage = error.error || errorMessage;
      } catch {
        // 如果响应不是 JSON，使用 HTTP 状态文本
        errorMessage = `HTTP ${response.status}: ${response.statusText}`;
      }
      throw new Error(errorMessage);
    }

    return response.json();
  }

  async getMessages(limit: number = 20, cursor?: number): Promise<MessageListResponse> {
    const params = new URLSearchParams({ limit: limit.toString() });
    if (cursor !== undefined) {
      params.append('cursor', cursor.toString());
    }

    const response = await fetch(`${this.baseURL}/messages?${params}`);

    if (!response.ok) {
      throw new Error('Failed to fetch messages');
    }

    return response.json();
  }

  async hugMessage(id: number): Promise<Message> {
    const response = await fetch(`${this.baseURL}/messages/${id}/hug`, {
      method: 'POST',
    });

    if (!response.ok) {
      let errorMessage = 'Failed to hug message';
      try {
        const error = await response.json();
        errorMessage = error.error || errorMessage;
      } catch {
        errorMessage = `HTTP ${response.status}: ${response.statusText}`;
      }
      throw new Error(errorMessage);
    }

    return response.json();
  }
}

export const apiClient = new APIClient(API_BASE_URL);
