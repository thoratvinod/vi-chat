// src/components/LeftSidebar.tsx
import React from 'react';

interface LeftSidebarProps {
  onSelectOption: (option: string) => void;
}

const LeftSidebar: React.FC<LeftSidebarProps> = ({ onSelectOption }) => {
  return (
    <div className="left-sidebar">
      <button onClick={() => onSelectOption('DirectMessage')}>Direct Message</button>
      <button onClick={() => onSelectOption('Group')}>Group</button>
    </div>
  );
};

export default LeftSidebar;
