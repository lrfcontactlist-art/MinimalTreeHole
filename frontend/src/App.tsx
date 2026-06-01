import { useState } from 'react';
import { MessageForm } from './components/MessageForm';
import { MessageList } from './components/MessageList';
import { apiClient } from './api/client';

function App() {
  const [refreshTrigger, setRefreshTrigger] = useState(0);

  const handleSubmit = async (content: string) => {
    await apiClient.createMessage({ content });
    setRefreshTrigger((prev) => prev + 1);
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <div className="container mx-auto px-4 py-8 max-w-3xl">
        <header className="text-center mb-8">
          <h1 className="text-4xl font-bold text-gray-800 mb-2">树洞</h1>
          <p className="text-gray-600">匿名分享你的心情</p>
        </header>

        <div className="mb-8">
          <MessageForm onSubmit={handleSubmit} />
        </div>

        <MessageList refreshTrigger={refreshTrigger} />
      </div>
    </div>
  );
}

export default App;
