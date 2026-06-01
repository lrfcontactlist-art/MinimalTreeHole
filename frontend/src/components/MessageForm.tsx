import React, { useState } from 'react';

interface MessageFormProps {
  onSubmit: (content: string) => Promise<void>;
}

export const MessageForm: React.FC<MessageFormProps> = ({ onSubmit }) => {
  const [content, setContent] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!content.trim()) {
      setError('请输入内容');
      return;
    }

    if (content.length > 500) {
      setError('内容不能超过500字');
      return;
    }

    setIsSubmitting(true);
    setError('');

    try {
      await onSubmit(content);
      setContent('');
    } catch (err) {
      setError(err instanceof Error ? err.message : '发布失败，请重试');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="bg-white rounded-lg shadow-md p-6">
      <textarea
        value={content}
        onChange={(e) => setContent(e.target.value)}
        placeholder="说点什么吧... (最多500字)"
        className="w-full h-32 p-4 border border-gray-300 rounded-lg resize-none focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        disabled={isSubmitting}
      />
      <div className="flex items-center justify-between mt-4">
        <span className={`text-sm ${content.length > 500 ? 'text-red-500' : 'text-gray-500'}`}>
          {content.length} / 500
        </span>
        <button
          type="submit"
          disabled={isSubmitting || !content.trim() || content.length > 500}
          className="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors"
        >
          {isSubmitting ? '发布中...' : '发布'}
        </button>
      </div>
      {error && (
        <p className="mt-2 text-sm text-red-500">{error}</p>
      )}
    </form>
  );
};
