import React from 'react';
import { Message } from '../types';

interface MessageCardProps {
  message: Message;
  onHug: (id: number) => void;
}

export const MessageCard: React.FC<MessageCardProps> = ({ message, onHug }) => {
  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    const now = new Date();
    const diff = now.getTime() - date.getTime();
    const minutes = Math.floor(diff / 60000);
    const hours = Math.floor(diff / 3600000);
    const days = Math.floor(diff / 86400000);

    if (minutes < 1) return '刚刚';
    if (minutes < 60) return `${minutes}分钟前`;
    if (hours < 24) return `${hours}小时前`;
    if (days < 7) return `${days}天前`;
    return date.toLocaleDateString('zh-CN');
  };

  return (
    <div className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow">
      <p className="text-gray-800 text-lg leading-relaxed mb-4 whitespace-pre-wrap">
        {message.content}
      </p>
      <div className="flex items-center justify-between text-sm text-gray-500">
        <span>{formatDate(message.created_at)}</span>
        <button
          onClick={() => onHug(message.id)}
          className="flex items-center gap-2 px-4 py-2 bg-pink-50 hover:bg-pink-100 text-pink-600 rounded-full transition-colors"
        >
          <span className="text-xl">🤗</span>
          <span className="font-medium">{message.hug_count}</span>
        </button>
      </div>
    </div>
  );
};
