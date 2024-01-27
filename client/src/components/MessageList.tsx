// src/components/MessageList.tsx
import React from 'react';

interface MessageListProps {
  messages: { id: number; sender: string; text: string }[];
}

const MessageList: React.FC<MessageListProps> = ({ messages }) => {
  return (
    <div className="message-list">
      {messages.map(message => (
        <div key={message.id}>
          <strong>{message.sender}</strong>: {message.text}
        </div>
      ))}
    </div>
  );
};

export default MessageList;
