import React, { useState, useEffect } from 'react';
import { Message } from '../types';
import { MessageCard } from './MessageCard';
import { apiClient } from '../api/client';

interface MessageListProps {
  refreshTrigger: number;
}

export const MessageList: React.FC<MessageListProps> = ({ refreshTrigger }) => {
  const [messages, setMessages] = useState<Message[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [nextCursor, setNextCursor] = useState<number | undefined>();
  const [loadingMore, setLoadingMore] = useState(false);

  const loadMessages = async (cursor?: number) => {
    try {
      const response = await apiClient.getMessages(20, cursor);
      if (cursor) {
        setMessages((prev) => [...prev, ...response.messages]);
      } else {
        setMessages(response.messages);
      }
      setNextCursor(response.next_cursor);
      setError('');
    } catch (err) {
      setError('加载失败，请刷新重试');
    } finally {
      setLoading(false);
      setLoadingMore(false);
    }
  };

  useEffect(() => {
    setLoading(true);
    loadMessages();
  }, [refreshTrigger]);

  const handleLoadMore = () => {
    if (nextCursor && !loadingMore) {
      setLoadingMore(true);
      loadMessages(nextCursor);
    }
  };

  const handleHug = async (id: number) => {
    try {
      const updatedMessage = await apiClient.hugMessage(id);
      setMessages((prev) =>
        prev.map((msg) => (msg.id === id ? updatedMessage : msg))
      );
    } catch (err) {
      console.error('Hug failed:', err);
    }
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center py-12">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="text-center py-12">
        <p className="text-red-500">{error}</p>
      </div>
    );
  }

  if (messages.length === 0) {
    return (
      <div className="text-center py-12">
        <p className="text-gray-500 text-lg">还没有留言，来发布第一条吧！</p>
      </div>
    );
  }

  return (
    <div className="space-y-4">
      {messages.map((message) => (
        <MessageCard key={message.id} message={message} onHug={handleHug} />
      ))}
      {nextCursor && (
        <div className="text-center py-4">
          <button
            onClick={handleLoadMore}
            disabled={loadingMore}
            className="px-6 py-2 bg-gray-100 hover:bg-gray-200 text-gray-700 rounded-lg disabled:opacity-50 transition-colors"
          >
            {loadingMore ? '加载中...' : '加载更多'}
          </button>
        </div>
      )}
    </div>
  );
};
