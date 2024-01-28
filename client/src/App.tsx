// src/App.tsx
import React, { useState } from 'react';
import LeftSidebar from './components/LeftSidebar';
import ContactList from './components/ContactList';
import MessageList from './components/MessageList';

const App: React.FC = () => {
  const [selectedOption, setSelectedOption] = useState<string | null>(null);
  const [selectedContact, setSelectedContact] = useState<{ id: number; name: string } | null>(null);

  const contacts = [
    { id: 1, name: 'User 1' },
    { id: 2, name: 'User 2' },
    // Add more contacts as needed
  ];

  const messages = [
    { id: 1, sender: 'User 1', text: 'Hello!' },
    { id: 2, sender: 'User 2', text: 'Hi there!' },
    // Add more messages as needed
  ];

  const handleSelectOption = (option: string) => {
    setSelectedOption(option);
    setSelectedContact(null);
  };

  const handleSelectContact = (contact: { id: number; name: string }) => {
    setSelectedContact(contact);
  };

  return (
    <div className="app">
      <LeftSidebar onSelectOption={handleSelectOption} />
      {selectedOption === 'DirectMessage' && (
        <ContactList contacts={contacts} onSelectContact={handleSelectContact} />
      )}
      {selectedContact && <MessageList messages={messages} />}
    </div>
  );
};

export default App;
