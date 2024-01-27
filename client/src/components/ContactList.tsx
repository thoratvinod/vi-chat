// src/components/ContactList.tsx
import React from 'react';

interface ContactListProps {
  contacts: { id: number; name: string }[];
  onSelectContact: (contact: { id: number; name: string }) => void;
}

const ContactList: React.FC<ContactListProps> = ({ contacts, onSelectContact }) => {
  return (
    <div className="contact-list">
      {contacts.map(contact => (
        <div key={contact.id} onClick={() => onSelectContact(contact)}>
          {contact.name}
        </div>
      ))}
    </div>
  );
};

export default ContactList;
